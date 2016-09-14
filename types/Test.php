<?php
/**
 * Auto generated from api.proto at 2016-09-12 04:24:17
 *
 * types package
 */

/**
 * Test message
 */
class Test extends \ProtobufMessage
{
    /* Field index constants */
    const A = 1;
    const B = 2;
    const C = 3;

    /* @var array Field descriptors */
    protected static $fields = array(
        self::A => array(
            'name' => 'A',
            'required' => false,
            'type' => \ProtobufMessage::PB_TYPE_INT,
        ),
        self::B => array(
            'name' => 'B',
            'required' => false,
            'type' => \ProtobufMessage::PB_TYPE_INT,
        ),
        self::C => array(
            'name' => 'C',
            'required' => false,
            'type' => \ProtobufMessage::PB_TYPE_INT,
        ),
    );

    /**
     * Constructs new message container and clears its internal state
     */
    public function __construct()
    {
        $this->reset();
    }

    /**
     * Clears message values and sets default ones
     *
     * @return null
     */
    public function reset()
    {
        $this->values[self::A] = null;
        $this->values[self::B] = null;
        $this->values[self::C] = null;
    }

    /**
     * Returns field descriptors
     *
     * @return array
     */
    public function fields()
    {
        return self::$fields;
    }

    /**
     * Sets value of 'A' property
     *
     * @param integer $value Property value
     *
     * @return null
     */
    public function setA($value)
    {
        return $this->set(self::A, $value);
    }

    /**
     * Returns value of 'A' property
     *
     * @return integer
     */
    public function getA()
    {
        $value = $this->get(self::A);
        return $value === null ? (integer)$value : $value;
    }

    /**
     * Sets value of 'B' property
     *
     * @param integer $value Property value
     *
     * @return null
     */
    public function setB($value)
    {
        return $this->set(self::B, $value);
    }

    /**
     * Returns value of 'B' property
     *
     * @return integer
     */
    public function getB()
    {
        $value = $this->get(self::B);
        return $value === null ? (integer)$value : $value;
    }

    /**
     * Sets value of 'C' property
     *
     * @param integer $value Property value
     *
     * @return null
     */
    public function setC($value)
    {
        return $this->set(self::C, $value);
    }

    /**
     * Returns value of 'C' property
     *
     * @return integer
     */
    public function getC()
    {
        $value = $this->get(self::C);
        return $value === null ? (integer)$value : $value;
    }
}