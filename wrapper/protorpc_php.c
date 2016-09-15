/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) 1997-2016 The PHP Group                                |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

/* $Id$ */

#ifdef HAVE_CONFIG_H
#include "config.h"
#endif

#include "php.h"
#include "php_ini.h"
#include "ext/standard/info.h"
#include "php_protorpc_php.h"
#include "protorpc.h"

/* If you declare any globals in php_protorpc_php.h uncomment this:
ZEND_DECLARE_MODULE_GLOBALS(protorpc_php)
*/

/* True global resources - no need for thread safety here */
static int le_protorpc_php;

/* {{{ PHP_INI
 */
/* Remove comments and fill if you need to have entries in php.ini
PHP_INI_BEGIN()
    STD_PHP_INI_ENTRY("protorpc_php.global_value",      "42", PHP_INI_ALL, OnUpdateLong, global_value, zend_protorpc_php_globals, protorpc_php_globals)
    STD_PHP_INI_ENTRY("protorpc_php.global_string", "foobar", PHP_INI_ALL, OnUpdateString, global_string, zend_protorpc_php_globals, protorpc_php_globals)
PHP_INI_END()
*/
/* }}} */

/* Remove the following function when you have successfully modified config.m4
   so that your module can be compiled into PHP, it exists only for testing
   purposes. */

/* Every user-visible function in PHP should document itself in the source */
/* {{{ proto string confirm_protorpc_php_compiled(string arg)
   Return a string to confirm that the module is compiled in */
PHP_FUNCTION(confirm_protorpc_php_compiled)
{
	char *arg = NULL;
	size_t arg_len, len;
	zend_string *strg;

	if (zend_parse_parameters(ZEND_NUM_ARGS(), "s", &arg, &arg_len) == FAILURE) {
		return;
	}

	strg = strpprintf(0, "Congratulations! You have successfully modified ext/%.78s/config.m4. Module %.78s is now compiled into PHP.", "protorpc_php", arg);

	RETURN_STR(strg);
}
/* }}} */
/* The previous line is meant for vim and emacs, so it can correctly fold and
   unfold functions in source code. See the corresponding marks just before
   function definition, where the functions purpose is also documented. Please
   follow this convention for the convenience of others editing your code.
*/


PHP_FUNCTION(protorpc)
{
	// 获取参数
	char *addr_char = NULL;
	size_t addr_len;
	long max, time_out;
	if (zend_parse_parameters(3, "sll", &addr_char, &addr_len,&max,&time_out) == FAILURE) {
		return;
	}

	// 处理请求
	GoString addr = {addr_char,strlen(addr_char)};
	GoString reply = Protorpc(addr,max,time_out);
	
	// 返回结果
	zend_string *strg;
	strg = zend_string_init(reply.p,strlen(reply.p),0);
	RETURN_STR(strg);
}

PHP_FUNCTION(protorpc_call)
{
	// 获取参数
	char *addr_char = NULL,*method_char = NULL,*data_char = NULL;
	size_t addr_len,method_len,data_len;
	if (zend_parse_parameters(3, "sss", &addr_char, &addr_len,&method_char, &method_len,&data_char, &data_len) == FAILURE) {
		return;
	}

	// 处理请求
	GoString addr = {addr_char,strlen(addr_char)},method = {method_char,strlen(method_char)},data = {data_char,strlen(data_char)};
	GoString reply = ProtorpcCall(addr,method,data);

	// 返回结果
	zend_string *strg;
	strg = zend_string_init(reply.p,strlen(reply.p),0);

	RETURN_STR(strg);
}


/* {{{ php_protorpc_php_init_globals
 */
/* Uncomment this function if you have INI entries
static void php_protorpc_php_init_globals(zend_protorpc_php_globals *protorpc_php_globals)
{
	protorpc_php_globals->global_value = 0;
	protorpc_php_globals->global_string = NULL;
}
*/
/* }}} */

/* {{{ PHP_MINIT_FUNCTION
 */
PHP_MINIT_FUNCTION(protorpc_php)
{
	/* If you have INI entries, uncomment these lines
	REGISTER_INI_ENTRIES();
	*/
	return SUCCESS;
}
/* }}} */

/* {{{ PHP_MSHUTDOWN_FUNCTION
 */
PHP_MSHUTDOWN_FUNCTION(protorpc_php)
{
	/* uncomment this line if you have INI entries
	UNREGISTER_INI_ENTRIES();
	*/
	return SUCCESS;
}
/* }}} */

/* Remove if there's nothing to do at request start */
/* {{{ PHP_RINIT_FUNCTION
 */
PHP_RINIT_FUNCTION(protorpc_php)
{
#if defined(COMPILE_DL_PROTORPC_PHP) && defined(ZTS)
	ZEND_TSRMLS_CACHE_UPDATE();
#endif
	return SUCCESS;
}
/* }}} */

/* Remove if there's nothing to do at request end */
/* {{{ PHP_RSHUTDOWN_FUNCTION
 */
PHP_RSHUTDOWN_FUNCTION(protorpc_php)
{
	return SUCCESS;
}
/* }}} */

/* {{{ PHP_MINFO_FUNCTION
 */
PHP_MINFO_FUNCTION(protorpc_php)
{
	php_info_print_table_start();
	php_info_print_table_header(2, "protorpc_php support", "enabled");
	php_info_print_table_end();

	/* Remove comments if you have entries in php.ini
	DISPLAY_INI_ENTRIES();
	*/
}
/* }}} */

/* {{{ protorpc_php_functions[]
 *
 * Every user visible function must have an entry in protorpc_php_functions[].
 */
const zend_function_entry protorpc_php_functions[] = {
	PHP_FE(protorpc,	NULL)	
	PHP_FE(protorpc_call,	NULL)	
	PHP_FE(confirm_protorpc_php_compiled,	NULL)		/* For testing, remove later. */
	PHP_FE_END	/* Must be the last line in protorpc_php_functions[] */
};
/* }}} */

/* {{{ protorpc_php_module_entry
 */
zend_module_entry protorpc_php_module_entry = {
	STANDARD_MODULE_HEADER,
	"protorpc_php",
	protorpc_php_functions,
	PHP_MINIT(protorpc_php),
	PHP_MSHUTDOWN(protorpc_php),
	PHP_RINIT(protorpc_php),		/* Replace with NULL if there's nothing to do at request start */
	PHP_RSHUTDOWN(protorpc_php),	/* Replace with NULL if there's nothing to do at request end */
	PHP_MINFO(protorpc_php),
	PHP_PROTORPC_PHP_VERSION,
	STANDARD_MODULE_PROPERTIES
};
/* }}} */

#ifdef COMPILE_DL_PROTORPC_PHP
#ifdef ZTS
ZEND_TSRMLS_CACHE_DEFINE()
#endif
ZEND_GET_MODULE(protorpc_php)
#endif

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: noet sw=4 ts=4 fdm=marker
 * vim<600: noet sw=4 ts=4
 */
