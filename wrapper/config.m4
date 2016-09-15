PHP_ARG_WITH(protorpc_php, for protorpc_php support,
Make sure that the comment is aligned:
[  --with-protorpc_php             Include protorpc_php support])

if test "$PHP_PROTORPC_PHP" != "no"; then

  PHP_SUBST(PROTORPC_PHP_SHARED_LIBADD)

  PHP_ADD_LIBRARY_WITH_PATH(protorpc,., PROTORPC_PHP_SHARED_LIBADD)

  PHP_NEW_EXTENSION(protorpc_php, protorpc_php.c, $ext_shared,, -DZEND_ENABLE_STATIC_TSRMLS_CACHE=1)
fi
