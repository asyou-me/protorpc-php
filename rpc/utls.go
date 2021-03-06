package rpc

import (
	"fmt"
	"hash/crc32"
	"io"

	types "github.com/asyou-me/protorpc/types"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
)

// 写入请求数据到 io.Writer
func writeRequest(c *clientCodec, id uint64, method string, pbRequest []byte) error {
	w := c.w

	// compress serialized proto data
	compressedPbRequest := snappy.Encode(nil, pbRequest)

	// generate header
	header := &types.RequestHeader{
		Id:                         id,
		Method:                     method,
		RawRequestLen:              uint32(len(pbRequest)),
		SnappyCompressedRequestLen: uint32(len(compressedPbRequest)),
		Checksum:                   crc32.ChecksumIEEE(compressedPbRequest),
		Token:                      c.Token,
	}

	// check header size
	pbHeader, err := proto.Marshal(header)
	if err != err {
		return err
	}
	if len(pbHeader) > int(types.Const_MAX_REQUEST_HEADER_LEN) {
		return fmt.Errorf("protorpc.writeRequest: header larger than max_header_len: %d. ", len(pbHeader))
	}

	// send header (more)
	if err := sendFrame(w, pbHeader); err != nil {
		return err
	}

	// send body (end)
	if err := sendFrame(w, compressedPbRequest); err != nil {
		return err
	}

	return nil
}

// 读取请求中的头信息
func readRequestHeader(r io.Reader, header *types.RequestHeader) (err error) {
	// recv header (more)
	pbHeader, err := recvFrame(r)
	if err != nil {
		return err
	}

	// Marshal Header
	err = proto.Unmarshal(pbHeader, header)
	if err != nil {
		return err
	}

	return nil
}

// 读取请求包内容
func readRequestBody(r io.Reader, header *types.RequestHeader, request proto.Message) error {
	// recv body (end)
	compressedPbRequest, err := recvFrame(r)
	if err != nil {
		return err
	}

	// checksum
	if crc32.ChecksumIEEE(compressedPbRequest) != header.Checksum {
		return fmt.Errorf("protorpc.readRequestBody: unexpected checksum. ")
	}

	// decode the compressed data
	pbRequest, err := snappy.Decode(nil, compressedPbRequest)
	if err != nil {
		return err
	}
	// check types header: rawMsgLen
	if uint32(len(pbRequest)) != header.RawRequestLen {
		return fmt.Errorf("protorpc.readRequestBody: Unexcpeted header.RawRequestLen. ")
	}

	// Unmarshal to proto message
	if request != nil {
		err = proto.Unmarshal(pbRequest, request)
		if err != nil {
			return err
		}
	}

	return nil
}

// 写入结果到客户端
func writeResponse(w io.Writer, id uint64, serr string, response proto.Message) (err error) {
	// clean response if error
	if serr != "" {
		response = nil
	}

	// marshal response
	pbResponse := []byte{}
	if response != nil {
		pbResponse, err = proto.Marshal(response)
		if err != nil {
			return err
		}
	}

	// compress serialized proto data
	compressedPbResponse := snappy.Encode(nil, pbResponse)

	// generate header
	header := &types.ResponseHeader{
		Id:                          id,
		Error:                       serr,
		RawResponseLen:              uint32(len(pbResponse)),
		SnappyCompressedResponseLen: uint32(len(compressedPbResponse)),
		Checksum:                    crc32.ChecksumIEEE(compressedPbResponse),
	}

	// check header size
	pbHeader, err := proto.Marshal(header)
	if err != err {
		return
	}

	// send header (more)
	if err = sendFrame(w, pbHeader); err != nil {
		return
	}

	// send body (end)
	if err = sendFrame(w, compressedPbResponse); err != nil {
		return
	}

	return nil
}

// 读取结果的头信息
func readResponseHeader(r io.Reader, header *types.ResponseHeader) error {
	// recv header (more)
	pbHeader, err := recvFrame(r)
	if err != nil {
		return err
	}

	// Marshal Header
	err = proto.Unmarshal(pbHeader, header)
	if err != nil {
		return err
	}

	return nil
}

// 读取结果的内容
func readResponseBody(r io.Reader, header *types.ResponseHeader, response *[]byte) error {
	// recv body (end)
	compressedPbResponse, err := recvFrame(r)
	if err != nil {
		return err
	}

	// checksum
	if crc32.ChecksumIEEE(compressedPbResponse) != header.Checksum {
		return fmt.Errorf("protorpc.readResponseBody: unexpected checksum. ")
	}

	// decode the compressed data
	pbResponse, err := snappy.Decode(nil, compressedPbResponse)
	if err != nil {
		return err
	}
	// check types header: rawMsgLen
	if uint32(len(pbResponse)) != header.RawResponseLen {
		return fmt.Errorf("protorpc.readResponseBody: Unexcpeted header.RawResponseLen. ")
	}

	// Unmarshal to proto message
	if response != nil {
		*response = pbResponse
	}

	return nil
}
