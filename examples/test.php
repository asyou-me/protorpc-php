<?php

require_once __DIR__."/../types/Test.php";

$test = new Test();
$test->setA(1);
$test->setB(2);
var_dump(strlen($test->serializeToString()));
var_dump(base64_encode($test->serializeToString()));

var_dump(protorpc("127.0.0.1:30015",10,11));

$replay = protorpc_call("127.0.0.1:30015","TestHandler.Test",$test->serializeToString());

var_dump($replay);

$test2 = new Test();
$test2->parseFromString($replay);

var_dump($test2);
