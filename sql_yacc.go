// CAUTION: Generated file - DO NOT EDIT.

package sql

import __yyfmt__ "fmt"

import ()

type MySQLSymType struct {
	yys       int
	bytes     []byte
	statement Statement
	table     Table
}

type MySQLXError struct {
	state, xsym int
}

const (
	MySQLDefault                      = 58010
	MySQLEofCode                      = 57344
	ABORT_SYM                         = 57346
	ACCESSIBLE_SYM                    = 57347
	ACCOUNT_SYM                       = 57348
	ACTION                            = 57349
	ADD                               = 57350
	ADDDATE_SYM                       = 57351
	ADMIN_SYM                         = 57998
	AFTER_SYM                         = 57352
	AGAINST                           = 57353
	AGGREGATE_SYM                     = 57354
	ALGORITHM_SYM                     = 57355
	ALL                               = 57356
	ALTER                             = 57357
	ALWAYS_SYM                        = 57358
	ANALYSE_SYM                       = 57359
	ANALYZE_SYM                       = 57360
	AND_AND_SYM                       = 57361
	AND_SYM                           = 57362
	ANY_SYM                           = 57363
	AS                                = 57364
	ASC                               = 57365
	ASCII_SYM                         = 57366
	ASENSITIVE_SYM                    = 57367
	AT_SYM                            = 57368
	AUTOEXTEND_SIZE_SYM               = 57369
	AUTO_INC                          = 57370
	AVG_ROW_LENGTH                    = 57371
	AVG_SYM                           = 57372
	BACKUP_SYM                        = 57373
	BEFORE_SYM                        = 57374
	BEGIN_SYM                         = 57375
	BETWEEN_SYM                       = 57376
	BIGINT_SYM                        = 57377
	BINARY_SYM                        = 57378
	BINLOG_SYM                        = 57379
	BIN_NUM                           = 57380
	BIT_AND                           = 57381
	BIT_OR                            = 57382
	BIT_SYM                           = 57383
	BIT_XOR                           = 57384
	BLOB_SYM                          = 57385
	BLOCK_SYM                         = 57386
	BOOLEAN_SYM                       = 57387
	BOOL_SYM                          = 57388
	BOTH                              = 57389
	BTREE_SYM                         = 57390
	BY                                = 57391
	BYTE_SYM                          = 57392
	CACHE_SYM                         = 57393
	CALL_SYM                          = 57394
	CASCADE                           = 57395
	CASCADED                          = 57396
	CASE_SYM                          = 57397
	CAST_SYM                          = 57398
	CATALOG_NAME_SYM                  = 57399
	CHAIN_SYM                         = 57400
	CHANGE                            = 57401
	CHANGED                           = 57402
	CHANNEL_SYM                       = 57403
	CHARSET                           = 57404
	CHAR_SYM                          = 57405
	CHECKSUM_SYM                      = 57406
	CHECK_SYM                         = 57407
	CIPHER_SYM                        = 57408
	CLASS_ORIGIN_SYM                  = 57409
	CLIENT_SYM                        = 57410
	CLOSE_SYM                         = 57411
	COALESCE                          = 57412
	CODE_SYM                          = 57413
	COLLATE_SYM                       = 57414
	COLLATION_SYM                     = 57415
	COLUMNS                           = 57416
	COLUMN_FORMAT_SYM                 = 57418
	COLUMN_NAME_SYM                   = 57419
	COLUMN_SYM                        = 57417
	COMMENT_SYM                       = 57420
	COMMITTED_SYM                     = 57421
	COMMIT_SYM                        = 57422
	COMPACT_SYM                       = 57423
	COMPLETION_SYM                    = 57424
	COMPONENT_SYM                     = 58002
	COMPRESSED_SYM                    = 57425
	COMPRESSION_SYM                   = 57426
	CONCURRENT                        = 57428
	CONDITIONLESS_JOIN                = 58006
	CONDITION_SYM                     = 57429
	CONNECTION_SYM                    = 57430
	CONSISTENT_SYM                    = 57431
	CONSTRAINT                        = 57432
	CONSTRAINT_CATALOG_SYM            = 57433
	CONSTRAINT_NAME_SYM               = 57434
	CONSTRAINT_SCHEMA_SYM             = 57435
	CONTAINS_SYM                      = 57436
	CONTEXT_SYM                       = 57437
	CONTINUE_SYM                      = 57438
	CONVERT_SYM                       = 57439
	COUNT_SYM                         = 57440
	CPU_SYM                           = 57441
	CREATE                            = 57442
	CROSS                             = 57443
	CUBE_SYM                          = 57444
	CURDATE                           = 57445
	CURRENT_SYM                       = 57446
	CURRENT_USER                      = 57447
	CURSOR_NAME_SYM                   = 57449
	CURSOR_SYM                        = 57448
	CURTIME                           = 57450
	DATABASE                          = 57451
	DATABASES                         = 57452
	DATAFILE_SYM                      = 57453
	DATA_SYM                          = 57454
	DATETIME_SYM                      = 57455
	DATE_ADD_INTERVAL                 = 57456
	DATE_SUB_INTERVAL                 = 57457
	DATE_SYM                          = 57458
	DAY_HOUR_SYM                      = 57459
	DAY_MICROSECOND_SYM               = 57460
	DAY_MINUTE_SYM                    = 57461
	DAY_SECOND_SYM                    = 57462
	DAY_SYM                           = 57463
	DEALLOCATE_SYM                    = 57464
	DECIMAL_NUM                       = 57465
	DECIMAL_SYM                       = 57466
	DECLARE_SYM                       = 57467
	DEFAULT_AUTH_SYM                  = 57469
	DEFAULT_SYM                       = 57468
	DEFINER_SYM                       = 57470
	DELAYED_SYM                       = 57471
	DELAY_KEY_WRITE_SYM               = 57472
	DELETE_SYM                        = 57473
	DESC                              = 57474
	DESCRIBE                          = 57475
	DES_KEY_FILE                      = 57476
	DETERMINISTIC_SYM                 = 57477
	DIAGNOSTICS_SYM                   = 57478
	DIRECTORY_SYM                     = 57479
	DISABLE_SYM                       = 57480
	DISCARD                           = 57481
	DISK_SYM                          = 57482
	DISTINCT                          = 57483
	DIV_SYM                           = 57484
	DOUBLE_SYM                        = 57485
	DO_SYM                            = 57486
	DROP                              = 57487
	DUAL_SYM                          = 57488
	DUMPFILE                          = 57489
	DUPLICATE_SYM                     = 57490
	DYNAMIC_SYM                       = 57491
	EACH_SYM                          = 57492
	ELSE                              = 57493
	ELSEIF_SYM                        = 57494
	EMPTY_FROM_CLAUSE                 = 58008
	ENABLE_SYM                        = 57495
	ENCLOSED                          = 57496
	ENCRYPTION_SYM                    = 57427
	END                               = 57497
	ENDS_SYM                          = 57498
	END_OF_INPUT                      = 57499
	ENGINES_SYM                       = 57500
	ENGINE_SYM                        = 57501
	ENUM_SYM                          = 57502
	EQ                                = 57503
	EQUAL_SYM                         = 57504
	ERRORS                            = 57506
	ERROR_SYM                         = 57505
	ESCAPED                           = 57507
	ESCAPE_SYM                        = 57508
	EVENTS_SYM                        = 57509
	EVENT_SYM                         = 57510
	EVERY_SYM                         = 57511
	EXCEPT_SYM                        = 58001
	EXCHANGE_SYM                      = 57512
	EXECUTE_SYM                       = 57513
	EXISTS                            = 57514
	EXIT_SYM                          = 57515
	EXPANSION_SYM                     = 57516
	EXPIRE_SYM                        = 57517
	EXPORT_SYM                        = 57518
	EXTENDED_SYM                      = 57519
	EXTENT_SIZE_SYM                   = 57520
	EXTRACT_SYM                       = 57521
	FALSE_SYM                         = 57522
	FAST_SYM                          = 57523
	FAULTS_SYM                        = 57524
	FETCH_SYM                         = 57525
	FILE_BLOCK_SIZE_SYM               = 57527
	FILE_SYM                          = 57526
	FILTER_SYM                        = 57528
	FIRST_SYM                         = 57529
	FIXED_SYM                         = 57530
	FLOAT_NUM                         = 57531
	FLOAT_SYM                         = 57532
	FLUSH_SYM                         = 57533
	FOLLOWS_SYM                       = 57534
	FORCE_SYM                         = 57535
	FOREIGN                           = 57536
	FORMAT_SYM                        = 57538
	FOR_SYM                           = 57537
	FOUND_SYM                         = 57539
	FROM                              = 57540
	FULL                              = 57541
	FULLTEXT_SYM                      = 57542
	FUNCTION_SYM                      = 57543
	GE                                = 57544
	GENERAL                           = 57545
	GENERATED                         = 57546
	GEOMETRYCOLLECTION_SYM            = 57548
	GEOMETRY_SYM                      = 57549
	GET_FORMAT                        = 57550
	GET_SYM                           = 57551
	GLOBAL_SYM                        = 57552
	GRAMMAR_SELECTOR_EXPR             = 58003
	GRAMMAR_SELECTOR_GCOL             = 58004
	GRAMMAR_SELECTOR_PART             = 58005
	GRANT                             = 57553
	GRANTS                            = 57554
	GROUP_CONCAT_SYM                  = 57556
	GROUP_REPLICATION                 = 57547
	GROUP_SYM                         = 57555
	GT_SYM                            = 57557
	HANDLER_SYM                       = 57558
	HASH_SYM                          = 57559
	HAVING                            = 57560
	HELP_SYM                          = 57561
	HEX_NUM                           = 57562
	HIGH_PRIORITY                     = 57563
	HOSTS_SYM                         = 57565
	HOST_SYM                          = 57564
	HOUR_MICROSECOND_SYM              = 57566
	HOUR_MINUTE_SYM                   = 57567
	HOUR_SECOND_SYM                   = 57568
	HOUR_SYM                          = 57569
	IDENT                             = 57570
	IDENTIFIED_SYM                    = 57571
	IDENT_QUOTED                      = 57572
	IF                                = 57573
	IGNORE_SERVER_IDS_SYM             = 57575
	IGNORE_SYM                        = 57574
	IMPORT                            = 57576
	INDEXES                           = 57577
	INDEX_SYM                         = 57578
	INFILE                            = 57579
	INITIAL_SIZE_SYM                  = 57580
	INNER_SYM                         = 57581
	INOUT_SYM                         = 57582
	INSENSITIVE_SYM                   = 57583
	INSERT_METHOD                     = 57585
	INSERT_SYM                        = 57584
	INSTALL_SYM                       = 57587
	INSTANCE_SYM                      = 57586
	INTERVAL_SYM                      = 57588
	INTO                              = 57589
	INT_SYM                           = 57590
	INVISIBLE_SYM                     = 57999
	INVOKER_SYM                       = 57591
	IN_SYM                            = 57592
	IO_AFTER_GTIDS                    = 57593
	IO_BEFORE_GTIDS                   = 57594
	IO_SYM                            = 57595
	IPC_SYM                           = 57596
	IS                                = 57597
	ISOLATION                         = 57598
	ISSUER_SYM                        = 57599
	ITERATE_SYM                       = 57600
	JOIN_SYM                          = 57601
	JSON_SEPARATOR_SYM                = 57602
	JSON_SYM                          = 57603
	JSON_UNQUOTED_SEPARATOR_SYM       = 57995
	KEYS                              = 57604
	KEY_BLOCK_SIZE                    = 57605
	KEY_SYM                           = 57606
	KILL_SYM                          = 57607
	LANGUAGE_SYM                      = 57608
	LAST_SYM                          = 57609
	LE                                = 57610
	LEADING                           = 57611
	LEAVES                            = 57612
	LEAVE_SYM                         = 57613
	LEFT                              = 57614
	LESS_SYM                          = 57615
	LEVEL_SYM                         = 57616
	LEX_HOSTNAME                      = 57617
	LIKE                              = 57618
	LIMIT                             = 57619
	LINEAR_SYM                        = 57620
	LINES                             = 57621
	LINESTRING_SYM                    = 57622
	LIST_SYM                          = 57623
	LOAD                              = 57624
	LOCAL_SYM                         = 57625
	LOCATOR_SYM                       = 57626
	LOCKS_SYM                         = 57627
	LOCK_SYM                          = 57628
	LOGFILE_SYM                       = 57629
	LOGS_SYM                          = 57630
	LONGBLOB_SYM                      = 57631
	LONGTEXT_SYM                      = 57632
	LONG_NUM                          = 57633
	LONG_SYM                          = 57634
	LOOP_SYM                          = 57635
	LOW_PRIORITY                      = 57636
	LT                                = 57637
	MASTER_AUTO_POSITION_SYM          = 57638
	MASTER_BIND_SYM                   = 57639
	MASTER_CONNECT_RETRY_SYM          = 57640
	MASTER_DELAY_SYM                  = 57641
	MASTER_HEARTBEAT_PERIOD_SYM       = 57661
	MASTER_HOST_SYM                   = 57642
	MASTER_LOG_FILE_SYM               = 57643
	MASTER_LOG_POS_SYM                = 57644
	MASTER_PASSWORD_SYM               = 57645
	MASTER_PORT_SYM                   = 57646
	MASTER_RETRY_COUNT_SYM            = 57647
	MASTER_SERVER_ID_SYM              = 57648
	MASTER_SSL_CAPATH_SYM             = 57649
	MASTER_SSL_CA_SYM                 = 57651
	MASTER_SSL_CERT_SYM               = 57652
	MASTER_SSL_CIPHER_SYM             = 57653
	MASTER_SSL_CRLPATH_SYM            = 57655
	MASTER_SSL_CRL_SYM                = 57654
	MASTER_SSL_KEY_SYM                = 57656
	MASTER_SSL_SYM                    = 57657
	MASTER_SSL_VERIFY_SERVER_CERT_SYM = 57658
	MASTER_SYM                        = 57659
	MASTER_TLS_VERSION_SYM            = 57650
	MASTER_USER_SYM                   = 57660
	MATCH                             = 57662
	MAX_CONNECTIONS_PER_HOUR          = 57663
	MAX_QUERIES_PER_HOUR              = 57664
	MAX_ROWS                          = 57665
	MAX_SIZE_SYM                      = 57666
	MAX_SYM                           = 57667
	MAX_UPDATES_PER_HOUR              = 57668
	MAX_USER_CONNECTIONS_SYM          = 57669
	MAX_VALUE_SYM                     = 57670
	MEDIUMBLOB_SYM                    = 57671
	MEDIUMINT_SYM                     = 57672
	MEDIUMTEXT_SYM                    = 57673
	MEDIUM_SYM                        = 57674
	MEMORY_SYM                        = 57675
	MERGE_SYM                         = 57676
	MESSAGE_TEXT_SYM                  = 57677
	MICROSECOND_SYM                   = 57678
	MIGRATE_SYM                       = 57679
	MINUTE_MICROSECOND_SYM            = 57680
	MINUTE_SECOND_SYM                 = 57681
	MINUTE_SYM                        = 57682
	MIN_ROWS                          = 57683
	MIN_SYM                           = 57684
	MODE_SYM                          = 57685
	MODIFIES_SYM                      = 57686
	MODIFY_SYM                        = 57687
	MOD_SYM                           = 57688
	MONTH_SYM                         = 57689
	MULTILINESTRING_SYM               = 57690
	MULTIPOINT_SYM                    = 57691
	MULTIPOLYGON_SYM                  = 57692
	MUTEX_SYM                         = 57693
	MYSQL_ERRNO_SYM                   = 57694
	NAMES_SYM                         = 57695
	NAME_SYM                          = 57696
	NATIONAL_SYM                      = 57697
	NATURAL                           = 57698
	NCHAR_STRING                      = 57699
	NCHAR_SYM                         = 57700
	NDBCLUSTER_SYM                    = 57701
	NE                                = 57702
	NEG                               = 57703
	NEVER_SYM                         = 57704
	NEW_SYM                           = 57705
	NEXT_SYM                          = 57706
	NODEGROUP_SYM                     = 57707
	NONE_SYM                          = 57708
	NOT2_SYM                          = 57709
	NOT_SYM                           = 57710
	NOW_SYM                           = 57711
	NO_SYM                            = 57712
	NO_WAIT_SYM                       = 57713
	NO_WRITE_TO_BINLOG                = 57714
	NULL_SYM                          = 57715
	NUM                               = 57716
	NUMBER_SYM                        = 57717
	NUMERIC_SYM                       = 57718
	NVARCHAR_SYM                      = 57719
	OBSOLETE_TOKEN_654                = 57742
	OBSOLETE_TOKEN_820                = 57908
	OFFSET_SYM                        = 57720
	ONE_SYM                           = 57722
	ONLY_SYM                          = 57723
	ON_SYM                            = 57721
	OPEN_SYM                          = 57724
	OPTIMIZE                          = 57725
	OPTIMIZER_COSTS_SYM               = 57726
	OPTION                            = 57728
	OPTIONALLY                        = 57729
	OPTIONS_SYM                       = 57727
	OR2_SYM                           = 57730
	ORDER_SYM                         = 57731
	OR_OR_SYM                         = 57732
	OR_SYM                            = 57733
	OUTER                             = 57734
	OUTFILE                           = 57735
	OUT_SYM                           = 57736
	OWNER_SYM                         = 57737
	PACK_KEYS_SYM                     = 57738
	PAGE_SYM                          = 57739
	PARAM_MARKER                      = 57740
	PARSER_SYM                        = 57741
	PARTIAL                           = 57743
	PARTITIONING_SYM                  = 57746
	PARTITIONS_SYM                    = 57745
	PARTITION_SYM                     = 57744
	PASSWORD                          = 57747
	PERSIST_SYM                       = 57996
	PHASE_SYM                         = 57748
	PLUGINS_SYM                       = 57751
	PLUGIN_DIR_SYM                    = 57749
	PLUGIN_SYM                        = 57750
	POINT_SYM                         = 57752
	POLYGON_SYM                       = 57753
	PORT_SYM                          = 57754
	POSITION_SYM                      = 57755
	PRECEDES_SYM                      = 57756
	PRECISION                         = 57757
	PREPARE_SYM                       = 57758
	PRESERVE_SYM                      = 57759
	PREV_SYM                          = 57760
	PRIMARY_SYM                       = 57761
	PRIVILEGES                        = 57762
	PROCEDURE_SYM                     = 57763
	PROCESS                           = 57764
	PROCESSLIST_SYM                   = 57765
	PROFILES_SYM                      = 57767
	PROFILE_SYM                       = 57766
	PROXY_SYM                         = 57768
	PURGE                             = 57769
	QUARTER_SYM                       = 57770
	QUERY_SYM                         = 57771
	QUICK                             = 57772
	RANGE_SYM                         = 57773
	READS_SYM                         = 57774
	READ_ONLY_SYM                     = 57775
	READ_SYM                          = 57776
	READ_WRITE_SYM                    = 57777
	REAL_SYM                          = 57778
	REBUILD_SYM                       = 57779
	RECOVER_SYM                       = 57780
	REDOFILE_SYM                      = 57781
	REDO_BUFFER_SIZE_SYM              = 57782
	REDUNDANT_SYM                     = 57783
	REFERENCES                        = 57784
	REGEXP                            = 57785
	RELAY                             = 57786
	RELAYLOG_SYM                      = 57787
	RELAY_LOG_FILE_SYM                = 57788
	RELAY_LOG_POS_SYM                 = 57789
	RELAY_THREAD                      = 57790
	RELEASE_SYM                       = 57791
	RELOAD                            = 57792
	REMOVE_SYM                        = 57793
	RENAME                            = 57794
	REORGANIZE_SYM                    = 57795
	REPAIR                            = 57796
	REPEATABLE_SYM                    = 57797
	REPEAT_SYM                        = 57798
	REPLACE_SYM                       = 57799
	REPLICATE_DO_DB                   = 57801
	REPLICATE_DO_TABLE                = 57803
	REPLICATE_IGNORE_DB               = 57802
	REPLICATE_IGNORE_TABLE            = 57804
	REPLICATE_REWRITE_DB              = 57807
	REPLICATE_WILD_DO_TABLE           = 57805
	REPLICATE_WILD_IGNORE_TABLE       = 57806
	REPLICATION                       = 57800
	REQUIRE_SYM                       = 57808
	RESET_SYM                         = 57809
	RESIGNAL_SYM                      = 57810
	RESOURCES                         = 57811
	RESTORE_SYM                       = 57812
	RESTRICT                          = 57813
	RESUME_SYM                        = 57814
	RETURNED_SQLSTATE_SYM             = 57815
	RETURNS_SYM                       = 57816
	RETURN_SYM                        = 57817
	REVERSE_SYM                       = 57818
	REVOKE                            = 57819
	RIGHT                             = 57820
	ROLE_SYM                          = 57997
	ROLLBACK_SYM                      = 57821
	ROLLUP_SYM                        = 57822
	ROTATE_SYM                        = 57823
	ROUTINE_SYM                       = 57824
	ROWS_SYM                          = 57825
	ROW_COUNT_SYM                     = 57828
	ROW_FORMAT_SYM                    = 57826
	ROW_SYM                           = 57827
	RTREE_SYM                         = 57829
	SAVEPOINT_SYM                     = 57830
	SCHEDULE_SYM                      = 57831
	SCHEMA_NAME_SYM                   = 57832
	SECOND_MICROSECOND_SYM            = 57833
	SECOND_SYM                        = 57834
	SECURITY_SYM                      = 57835
	SELECT_SYM                        = 57836
	SENSITIVE_SYM                     = 57837
	SEPARATOR_SYM                     = 57838
	SERIALIZABLE_SYM                  = 57839
	SERIAL_SYM                        = 57840
	SERVER_OPTIONS                    = 57843
	SERVER_SYM                        = 57842
	SESSION_SYM                       = 57841
	SET_SYM                           = 57844
	SET_VAR                           = 57845
	SHARE_SYM                         = 57846
	SHIFT_LEFT                        = 57847
	SHIFT_RIGHT                       = 57848
	SHOW                              = 57849
	SHUTDOWN                          = 57850
	SIGNAL_SYM                        = 57851
	SIGNED_SYM                        = 57852
	SIMPLE_SYM                        = 57853
	SLAVE                             = 57854
	SLOW                              = 57855
	SMALLINT_SYM                      = 57856
	SNAPSHOT_SYM                      = 57857
	SOCKET_SYM                        = 57858
	SONAME_SYM                        = 57859
	SOUNDS_SYM                        = 57860
	SOURCE_SYM                        = 57861
	SPATIAL_SYM                       = 57862
	SPECIFIC_SYM                      = 57863
	SQLEXCEPTION_SYM                  = 57864
	SQLSTATE_SYM                      = 57865
	SQLWARNING_SYM                    = 57866
	SQL_AFTER_GTIDS                   = 57867
	SQL_AFTER_MTS_GAPS                = 57868
	SQL_BEFORE_GTIDS                  = 57869
	SQL_BIG_RESULT                    = 57870
	SQL_BUFFER_RESULT                 = 57871
	SQL_CACHE_SYM                     = 57872
	SQL_CALC_FOUND_ROWS               = 57873
	SQL_NO_CACHE_SYM                  = 57874
	SQL_SMALL_RESULT                  = 57875
	SQL_SYM                           = 57876
	SQL_THREAD                        = 57877
	SSL_SYM                           = 57878
	STACKED_SYM                       = 57879
	STARTING                          = 57880
	STARTS_SYM                        = 57881
	START_SYM                         = 57882
	STATS_AUTO_RECALC_SYM             = 57883
	STATS_PERSISTENT_SYM              = 57884
	STATS_SAMPLE_PAGES_SYM            = 57885
	STATUS_SYM                        = 57886
	STDDEV_SAMP_SYM                   = 57887
	STD_SYM                           = 57888
	STOP_SYM                          = 57889
	STORAGE_SYM                       = 57890
	STORED_SYM                        = 57891
	STRAIGHT_JOIN                     = 57892
	STRING_SYM                        = 57893
	SUBCLASS_ORIGIN_SYM               = 57894
	SUBDATE_SYM                       = 57895
	SUBJECT_SYM                       = 57896
	SUBPARTITIONS_SYM                 = 57897
	SUBPARTITION_SYM                  = 57898
	SUBQUERY_AS_EXPR                  = 58007
	SUBSTRING                         = 57899
	SUM_SYM                           = 57900
	SUPER_SYM                         = 57901
	SUSPEND_SYM                       = 57902
	SWAPS_SYM                         = 57903
	SWITCHES_SYM                      = 57904
	SYSDATE                           = 57905
	TABLES                            = 57906
	TABLESPACE_SYM                    = 57907
	TABLE_CHECKSUM_SYM                = 57910
	TABLE_NAME_SYM                    = 57911
	TABLE_SYM                         = 57909
	TEMPORARY                         = 57912
	TEMPTABLE_SYM                     = 57913
	TERMINATED                        = 57914
	TEXT_STRING                       = 57915
	TEXT_SYM                          = 57916
	THAN_SYM                          = 57917
	THEN_SYM                          = 57918
	TIMESTAMP_ADD                     = 57920
	TIMESTAMP_DIFF                    = 57921
	TIMESTAMP_SYM                     = 57919
	TIME_SYM                          = 57922
	TINYBLOB_SYM                      = 57923
	TINYINT_SYM                       = 57924
	TINYTEXT_SYN                      = 57925
	TO_SYM                            = 57926
	TRAILING                          = 57927
	TRANSACTION_SYM                   = 57928
	TRIGGERS_SYM                      = 57929
	TRIGGER_SYM                       = 57930
	TRIM                              = 57931
	TRUE_SYM                          = 57932
	TRUNCATE_SYM                      = 57933
	TYPES_SYM                         = 57934
	TYPE_SYM                          = 57935
	UDF_RETURNS_SYM                   = 57936
	ULONGLONG_NUM                     = 57937
	UNCOMMITTED_SYM                   = 57938
	UNDEFINED_SYM                     = 57939
	UNDERSCORE_CHARSET                = 57940
	UNDOFILE_SYM                      = 57941
	UNDO_BUFFER_SIZE_SYM              = 57942
	UNDO_SYM                          = 57943
	UNICODE_SYM                       = 57944
	UNINSTALL_SYM                     = 57945
	UNION_SYM                         = 57946
	UNIQUE_SYM                        = 57947
	UNKNOWN_SYM                       = 57948
	UNLOCK_SYM                        = 57949
	UNSIGNED_SYM                      = 57950
	UNTIL_SYM                         = 57951
	UPDATE_SYM                        = 57952
	UPGRADE_SYM                       = 57953
	USAGE                             = 57954
	USER                              = 57955
	USE_FRM                           = 57956
	USE_SYM                           = 57957
	USING                             = 57958
	UTC_DATE_SYM                      = 57959
	UTC_TIMESTAMP_SYM                 = 57960
	UTC_TIME_SYM                      = 57961
	VALIDATION_SYM                    = 57962
	VALUES                            = 57963
	VALUE_SYM                         = 57964
	VARBINARY_SYM                     = 57965
	VARCHAR_SYM                       = 57966
	VARIABLES                         = 57967
	VARIANCE_SYM                      = 57968
	VARYING                           = 57969
	VAR_SAMP_SYM                      = 57970
	VIEW_SYM                          = 57971
	VIRTUAL_SYM                       = 57972
	VISIBLE_SYM                       = 58000
	WAIT_SYM                          = 57973
	WARNINGS                          = 57974
	WEEK_SYM                          = 57975
	WEIGHT_STRING_SYM                 = 57976
	WHEN_SYM                          = 57977
	WHERE                             = 57978
	WHILE_SYM                         = 57979
	WITH                              = 57980
	WITHOUT_SYM                       = 57983
	WITH_CUBE_SYM                     = 57981
	WITH_ROLLUP_SYM                   = 57982
	WORK_SYM                          = 57984
	WRAPPER_SYM                       = 57985
	WRITE_SYM                         = 57986
	X509_SYM                          = 57987
	XA_SYM                            = 57988
	XID_SYM                           = 57989
	XML_SYM                           = 57990
	XOR                               = 57991
	YEAR_MONTH_SYM                    = 57992
	YEAR_SYM                          = 57993
	ZEROFILL_SYM                      = 57994
	MySQLErrCode                      = 57345

	MySQLMaxDepth = 200
	MySQLTabOfs   = -399
)

var (
	MySQLXLAT = map[int]int{
		57499: 0,   // END_OF_INPUT (394x)
		59:    1,   // ';' (393x)
		46:    2,   // '.' (387x)
		57344: 3,   // $end (7x)
		57348: 4,   // ACCOUNT_SYM (3x)
		57349: 5,   // ACTION (3x)
		57351: 6,   // ADDDATE_SYM (3x)
		57352: 7,   // AFTER_SYM (3x)
		57353: 8,   // AGAINST (3x)
		57354: 9,   // AGGREGATE_SYM (3x)
		57355: 10,  // ALGORITHM_SYM (3x)
		57358: 11,  // ALWAYS_SYM (3x)
		57359: 12,  // ANALYSE_SYM (3x)
		57363: 13,  // ANY_SYM (3x)
		57366: 14,  // ASCII_SYM (3x)
		57368: 15,  // AT_SYM (3x)
		57370: 16,  // AUTO_INC (3x)
		57369: 17,  // AUTOEXTEND_SIZE_SYM (3x)
		57371: 18,  // AVG_ROW_LENGTH (3x)
		57372: 19,  // AVG_SYM (3x)
		57373: 20,  // BACKUP_SYM (3x)
		57375: 21,  // BEGIN_SYM (3x)
		57379: 22,  // BINLOG_SYM (3x)
		57383: 23,  // BIT_SYM (3x)
		57386: 24,  // BLOCK_SYM (3x)
		57388: 25,  // BOOL_SYM (3x)
		57387: 26,  // BOOLEAN_SYM (3x)
		57390: 27,  // BTREE_SYM (3x)
		57392: 28,  // BYTE_SYM (3x)
		57393: 29,  // CACHE_SYM (3x)
		57396: 30,  // CASCADED (3x)
		57399: 31,  // CATALOG_NAME_SYM (3x)
		57400: 32,  // CHAIN_SYM (3x)
		57402: 33,  // CHANGED (3x)
		57403: 34,  // CHANNEL_SYM (3x)
		57404: 35,  // CHARSET (3x)
		57406: 36,  // CHECKSUM_SYM (3x)
		57408: 37,  // CIPHER_SYM (3x)
		57409: 38,  // CLASS_ORIGIN_SYM (3x)
		57410: 39,  // CLIENT_SYM (3x)
		57411: 40,  // CLOSE_SYM (3x)
		57412: 41,  // COALESCE (3x)
		57413: 42,  // CODE_SYM (3x)
		57415: 43,  // COLLATION_SYM (3x)
		57418: 44,  // COLUMN_FORMAT_SYM (3x)
		57419: 45,  // COLUMN_NAME_SYM (3x)
		57416: 46,  // COLUMNS (3x)
		57420: 47,  // COMMENT_SYM (3x)
		57422: 48,  // COMMIT_SYM (3x)
		57421: 49,  // COMMITTED_SYM (3x)
		57423: 50,  // COMPACT_SYM (3x)
		57424: 51,  // COMPLETION_SYM (3x)
		58002: 52,  // COMPONENT_SYM (3x)
		57425: 53,  // COMPRESSED_SYM (3x)
		57426: 54,  // COMPRESSION_SYM (3x)
		57428: 55,  // CONCURRENT (3x)
		57430: 56,  // CONNECTION_SYM (3x)
		57431: 57,  // CONSISTENT_SYM (3x)
		57433: 58,  // CONSTRAINT_CATALOG_SYM (3x)
		57434: 59,  // CONSTRAINT_NAME_SYM (3x)
		57435: 60,  // CONSTRAINT_SCHEMA_SYM (3x)
		57436: 61,  // CONTAINS_SYM (3x)
		57437: 62,  // CONTEXT_SYM (3x)
		57441: 63,  // CPU_SYM (3x)
		57444: 64,  // CUBE_SYM (3x)
		57446: 65,  // CURRENT_SYM (3x)
		57449: 66,  // CURSOR_NAME_SYM (3x)
		57454: 67,  // DATA_SYM (3x)
		57453: 68,  // DATAFILE_SYM (3x)
		57458: 69,  // DATE_SYM (3x)
		57455: 70,  // DATETIME_SYM (3x)
		57463: 71,  // DAY_SYM (3x)
		57464: 72,  // DEALLOCATE_SYM (3x)
		57469: 73,  // DEFAULT_AUTH_SYM (3x)
		57470: 74,  // DEFINER_SYM (3x)
		57472: 75,  // DELAY_KEY_WRITE_SYM (3x)
		57476: 76,  // DES_KEY_FILE (3x)
		57478: 77,  // DIAGNOSTICS_SYM (3x)
		57479: 78,  // DIRECTORY_SYM (3x)
		57480: 79,  // DISABLE_SYM (3x)
		57481: 80,  // DISCARD (3x)
		57482: 81,  // DISK_SYM (3x)
		57486: 82,  // DO_SYM (3x)
		57489: 83,  // DUMPFILE (3x)
		57490: 84,  // DUPLICATE_SYM (3x)
		57491: 85,  // DYNAMIC_SYM (3x)
		57495: 86,  // ENABLE_SYM (3x)
		57427: 87,  // ENCRYPTION_SYM (3x)
		57497: 88,  // END (3x)
		57498: 89,  // ENDS_SYM (3x)
		57501: 90,  // ENGINE_SYM (3x)
		57500: 91,  // ENGINES_SYM (3x)
		57502: 92,  // ENUM_SYM (3x)
		57505: 93,  // ERROR_SYM (3x)
		57506: 94,  // ERRORS (3x)
		57508: 95,  // ESCAPE_SYM (3x)
		57510: 96,  // EVENT_SYM (3x)
		57509: 97,  // EVENTS_SYM (3x)
		57511: 98,  // EVERY_SYM (3x)
		57512: 99,  // EXCHANGE_SYM (3x)
		57513: 100, // EXECUTE_SYM (3x)
		57516: 101, // EXPANSION_SYM (3x)
		57517: 102, // EXPIRE_SYM (3x)
		57518: 103, // EXPORT_SYM (3x)
		57519: 104, // EXTENDED_SYM (3x)
		57520: 105, // EXTENT_SIZE_SYM (3x)
		57523: 106, // FAST_SYM (3x)
		57524: 107, // FAULTS_SYM (3x)
		57527: 108, // FILE_BLOCK_SIZE_SYM (3x)
		57526: 109, // FILE_SYM (3x)
		57528: 110, // FILTER_SYM (3x)
		57529: 111, // FIRST_SYM (3x)
		57530: 112, // FIXED_SYM (3x)
		57533: 113, // FLUSH_SYM (3x)
		57534: 114, // FOLLOWS_SYM (3x)
		57538: 115, // FORMAT_SYM (3x)
		57539: 116, // FOUND_SYM (3x)
		57541: 117, // FULL (3x)
		57543: 118, // FUNCTION_SYM (3x)
		57545: 119, // GENERAL (3x)
		57549: 120, // GEOMETRY_SYM (3x)
		57548: 121, // GEOMETRYCOLLECTION_SYM (3x)
		57550: 122, // GET_FORMAT (3x)
		57552: 123, // GLOBAL_SYM (3x)
		57554: 124, // GRANTS (3x)
		57547: 125, // GROUP_REPLICATION (3x)
		57558: 126, // HANDLER_SYM (3x)
		57559: 127, // HASH_SYM (3x)
		57561: 128, // HELP_SYM (3x)
		57564: 129, // HOST_SYM (3x)
		57565: 130, // HOSTS_SYM (3x)
		57569: 131, // HOUR_SYM (3x)
		58014: 132, // ident (3x)
		57570: 133, // IDENT (3x)
		58015: 134, // ident_keyword (3x)
		57572: 135, // IDENT_QUOTED (3x)
		58011: 136, // IDENT_sys (3x)
		57571: 137, // IDENTIFIED_SYM (3x)
		57575: 138, // IGNORE_SERVER_IDS_SYM (3x)
		57576: 139, // IMPORT (3x)
		57577: 140, // INDEXES (3x)
		57580: 141, // INITIAL_SIZE_SYM (3x)
		57585: 142, // INSERT_METHOD (3x)
		57587: 143, // INSTALL_SYM (3x)
		57586: 144, // INSTANCE_SYM (3x)
		57999: 145, // INVISIBLE_SYM (3x)
		57591: 146, // INVOKER_SYM (3x)
		57595: 147, // IO_SYM (3x)
		57596: 148, // IPC_SYM (3x)
		57598: 149, // ISOLATION (3x)
		57599: 150, // ISSUER_SYM (3x)
		57603: 151, // JSON_SYM (3x)
		57605: 152, // KEY_BLOCK_SIZE (3x)
		58017: 153, // label_keyword (3x)
		57608: 154, // LANGUAGE_SYM (3x)
		57609: 155, // LAST_SYM (3x)
		57612: 156, // LEAVES (3x)
		57615: 157, // LESS_SYM (3x)
		57616: 158, // LEVEL_SYM (3x)
		57622: 159, // LINESTRING_SYM (3x)
		57623: 160, // LIST_SYM (3x)
		57625: 161, // LOCAL_SYM (3x)
		57627: 162, // LOCKS_SYM (3x)
		57629: 163, // LOGFILE_SYM (3x)
		57630: 164, // LOGS_SYM (3x)
		57638: 165, // MASTER_AUTO_POSITION_SYM (3x)
		57640: 166, // MASTER_CONNECT_RETRY_SYM (3x)
		57641: 167, // MASTER_DELAY_SYM (3x)
		57661: 168, // MASTER_HEARTBEAT_PERIOD_SYM (3x)
		57642: 169, // MASTER_HOST_SYM (3x)
		57643: 170, // MASTER_LOG_FILE_SYM (3x)
		57644: 171, // MASTER_LOG_POS_SYM (3x)
		57645: 172, // MASTER_PASSWORD_SYM (3x)
		57646: 173, // MASTER_PORT_SYM (3x)
		57647: 174, // MASTER_RETRY_COUNT_SYM (3x)
		57648: 175, // MASTER_SERVER_ID_SYM (3x)
		57651: 176, // MASTER_SSL_CA_SYM (3x)
		57649: 177, // MASTER_SSL_CAPATH_SYM (3x)
		57652: 178, // MASTER_SSL_CERT_SYM (3x)
		57653: 179, // MASTER_SSL_CIPHER_SYM (3x)
		57654: 180, // MASTER_SSL_CRL_SYM (3x)
		57655: 181, // MASTER_SSL_CRLPATH_SYM (3x)
		57656: 182, // MASTER_SSL_KEY_SYM (3x)
		57657: 183, // MASTER_SSL_SYM (3x)
		57659: 184, // MASTER_SYM (3x)
		57650: 185, // MASTER_TLS_VERSION_SYM (3x)
		57660: 186, // MASTER_USER_SYM (3x)
		57663: 187, // MAX_CONNECTIONS_PER_HOUR (3x)
		57664: 188, // MAX_QUERIES_PER_HOUR (3x)
		57665: 189, // MAX_ROWS (3x)
		57666: 190, // MAX_SIZE_SYM (3x)
		57668: 191, // MAX_UPDATES_PER_HOUR (3x)
		57669: 192, // MAX_USER_CONNECTIONS_SYM (3x)
		57674: 193, // MEDIUM_SYM (3x)
		57675: 194, // MEMORY_SYM (3x)
		57676: 195, // MERGE_SYM (3x)
		57677: 196, // MESSAGE_TEXT_SYM (3x)
		57678: 197, // MICROSECOND_SYM (3x)
		57679: 198, // MIGRATE_SYM (3x)
		57683: 199, // MIN_ROWS (3x)
		57682: 200, // MINUTE_SYM (3x)
		57685: 201, // MODE_SYM (3x)
		57687: 202, // MODIFY_SYM (3x)
		57689: 203, // MONTH_SYM (3x)
		57690: 204, // MULTILINESTRING_SYM (3x)
		57691: 205, // MULTIPOINT_SYM (3x)
		57692: 206, // MULTIPOLYGON_SYM (3x)
		57693: 207, // MUTEX_SYM (3x)
		57694: 208, // MYSQL_ERRNO_SYM (3x)
		57696: 209, // NAME_SYM (3x)
		57695: 210, // NAMES_SYM (3x)
		57697: 211, // NATIONAL_SYM (3x)
		57700: 212, // NCHAR_SYM (3x)
		57701: 213, // NDBCLUSTER_SYM (3x)
		57704: 214, // NEVER_SYM (3x)
		57705: 215, // NEW_SYM (3x)
		57706: 216, // NEXT_SYM (3x)
		57712: 217, // NO_SYM (3x)
		57713: 218, // NO_WAIT_SYM (3x)
		57707: 219, // NODEGROUP_SYM (3x)
		57708: 220, // NONE_SYM (3x)
		57717: 221, // NUMBER_SYM (3x)
		57719: 222, // NVARCHAR_SYM (3x)
		57720: 223, // OFFSET_SYM (3x)
		57722: 224, // ONE_SYM (3x)
		57723: 225, // ONLY_SYM (3x)
		57724: 226, // OPEN_SYM (3x)
		57727: 227, // OPTIONS_SYM (3x)
		57737: 228, // OWNER_SYM (3x)
		57738: 229, // PACK_KEYS_SYM (3x)
		57739: 230, // PAGE_SYM (3x)
		57741: 231, // PARSER_SYM (3x)
		57743: 232, // PARTIAL (3x)
		57746: 233, // PARTITIONING_SYM (3x)
		57745: 234, // PARTITIONS_SYM (3x)
		57747: 235, // PASSWORD (3x)
		57748: 236, // PHASE_SYM (3x)
		57749: 237, // PLUGIN_DIR_SYM (3x)
		57750: 238, // PLUGIN_SYM (3x)
		57751: 239, // PLUGINS_SYM (3x)
		57752: 240, // POINT_SYM (3x)
		57753: 241, // POLYGON_SYM (3x)
		57754: 242, // PORT_SYM (3x)
		57756: 243, // PRECEDES_SYM (3x)
		57758: 244, // PREPARE_SYM (3x)
		57759: 245, // PRESERVE_SYM (3x)
		57760: 246, // PREV_SYM (3x)
		57762: 247, // PRIVILEGES (3x)
		57764: 248, // PROCESS (3x)
		57765: 249, // PROCESSLIST_SYM (3x)
		57766: 250, // PROFILE_SYM (3x)
		57767: 251, // PROFILES_SYM (3x)
		57768: 252, // PROXY_SYM (3x)
		57770: 253, // QUARTER_SYM (3x)
		57771: 254, // QUERY_SYM (3x)
		57772: 255, // QUICK (3x)
		57775: 256, // READ_ONLY_SYM (3x)
		57779: 257, // REBUILD_SYM (3x)
		57780: 258, // RECOVER_SYM (3x)
		57782: 259, // REDO_BUFFER_SIZE_SYM (3x)
		57781: 260, // REDOFILE_SYM (3x)
		57783: 261, // REDUNDANT_SYM (3x)
		57786: 262, // RELAY (3x)
		57788: 263, // RELAY_LOG_FILE_SYM (3x)
		57789: 264, // RELAY_LOG_POS_SYM (3x)
		57790: 265, // RELAY_THREAD (3x)
		57787: 266, // RELAYLOG_SYM (3x)
		57792: 267, // RELOAD (3x)
		57793: 268, // REMOVE_SYM (3x)
		57795: 269, // REORGANIZE_SYM (3x)
		57796: 270, // REPAIR (3x)
		57797: 271, // REPEATABLE_SYM (3x)
		57801: 272, // REPLICATE_DO_DB (3x)
		57803: 273, // REPLICATE_DO_TABLE (3x)
		57802: 274, // REPLICATE_IGNORE_DB (3x)
		57804: 275, // REPLICATE_IGNORE_TABLE (3x)
		57807: 276, // REPLICATE_REWRITE_DB (3x)
		57805: 277, // REPLICATE_WILD_DO_TABLE (3x)
		57806: 278, // REPLICATE_WILD_IGNORE_TABLE (3x)
		57800: 279, // REPLICATION (3x)
		57809: 280, // RESET_SYM (3x)
		57811: 281, // RESOURCES (3x)
		57812: 282, // RESTORE_SYM (3x)
		57814: 283, // RESUME_SYM (3x)
		57815: 284, // RETURNED_SQLSTATE_SYM (3x)
		57816: 285, // RETURNS_SYM (3x)
		57818: 286, // REVERSE_SYM (3x)
		58019: 287, // role_or_ident_keyword (3x)
		58020: 288, // role_or_label_keyword (3x)
		57821: 289, // ROLLBACK_SYM (3x)
		57822: 290, // ROLLUP_SYM (3x)
		57823: 291, // ROTATE_SYM (3x)
		57824: 292, // ROUTINE_SYM (3x)
		57828: 293, // ROW_COUNT_SYM (3x)
		57826: 294, // ROW_FORMAT_SYM (3x)
		57827: 295, // ROW_SYM (3x)
		57825: 296, // ROWS_SYM (3x)
		57829: 297, // RTREE_SYM (3x)
		57830: 298, // SAVEPOINT_SYM (3x)
		57831: 299, // SCHEDULE_SYM (3x)
		57832: 300, // SCHEMA_NAME_SYM (3x)
		57834: 301, // SECOND_SYM (3x)
		57835: 302, // SECURITY_SYM (3x)
		57840: 303, // SERIAL_SYM (3x)
		57839: 304, // SERIALIZABLE_SYM (3x)
		57842: 305, // SERVER_SYM (3x)
		57841: 306, // SESSION_SYM (3x)
		57846: 307, // SHARE_SYM (3x)
		57850: 308, // SHUTDOWN (3x)
		57852: 309, // SIGNED_SYM (3x)
		57853: 310, // SIMPLE_SYM (3x)
		57854: 311, // SLAVE (3x)
		57855: 312, // SLOW (3x)
		57857: 313, // SNAPSHOT_SYM (3x)
		57858: 314, // SOCKET_SYM (3x)
		57859: 315, // SONAME_SYM (3x)
		57860: 316, // SOUNDS_SYM (3x)
		57861: 317, // SOURCE_SYM (3x)
		57867: 318, // SQL_AFTER_GTIDS (3x)
		57868: 319, // SQL_AFTER_MTS_GAPS (3x)
		57869: 320, // SQL_BEFORE_GTIDS (3x)
		57871: 321, // SQL_BUFFER_RESULT (3x)
		57872: 322, // SQL_CACHE_SYM (3x)
		57874: 323, // SQL_NO_CACHE_SYM (3x)
		57877: 324, // SQL_THREAD (3x)
		57879: 325, // STACKED_SYM (3x)
		57882: 326, // START_SYM (3x)
		57881: 327, // STARTS_SYM (3x)
		57883: 328, // STATS_AUTO_RECALC_SYM (3x)
		57884: 329, // STATS_PERSISTENT_SYM (3x)
		57885: 330, // STATS_SAMPLE_PAGES_SYM (3x)
		57886: 331, // STATUS_SYM (3x)
		57889: 332, // STOP_SYM (3x)
		57890: 333, // STORAGE_SYM (3x)
		57893: 334, // STRING_SYM (3x)
		57894: 335, // SUBCLASS_ORIGIN_SYM (3x)
		57895: 336, // SUBDATE_SYM (3x)
		57896: 337, // SUBJECT_SYM (3x)
		57898: 338, // SUBPARTITION_SYM (3x)
		57897: 339, // SUBPARTITIONS_SYM (3x)
		57901: 340, // SUPER_SYM (3x)
		57902: 341, // SUSPEND_SYM (3x)
		57903: 342, // SWAPS_SYM (3x)
		57904: 343, // SWITCHES_SYM (3x)
		57910: 344, // TABLE_CHECKSUM_SYM (3x)
		57911: 345, // TABLE_NAME_SYM (3x)
		57906: 346, // TABLES (3x)
		57907: 347, // TABLESPACE_SYM (3x)
		57912: 348, // TEMPORARY (3x)
		57913: 349, // TEMPTABLE_SYM (3x)
		57916: 350, // TEXT_SYM (3x)
		57917: 351, // THAN_SYM (3x)
		57922: 352, // TIME_SYM (3x)
		57920: 353, // TIMESTAMP_ADD (3x)
		57921: 354, // TIMESTAMP_DIFF (3x)
		57919: 355, // TIMESTAMP_SYM (3x)
		57928: 356, // TRANSACTION_SYM (3x)
		57929: 357, // TRIGGERS_SYM (3x)
		57933: 358, // TRUNCATE_SYM (3x)
		57935: 359, // TYPE_SYM (3x)
		57934: 360, // TYPES_SYM (3x)
		57936: 361, // UDF_RETURNS_SYM (3x)
		57938: 362, // UNCOMMITTED_SYM (3x)
		57939: 363, // UNDEFINED_SYM (3x)
		57942: 364, // UNDO_BUFFER_SIZE_SYM (3x)
		57941: 365, // UNDOFILE_SYM (3x)
		57944: 366, // UNICODE_SYM (3x)
		57945: 367, // UNINSTALL_SYM (3x)
		57948: 368, // UNKNOWN_SYM (3x)
		57951: 369, // UNTIL_SYM (3x)
		57953: 370, // UPGRADE_SYM (3x)
		57956: 371, // USE_FRM (3x)
		57955: 372, // USER (3x)
		57962: 373, // VALIDATION_SYM (3x)
		57964: 374, // VALUE_SYM (3x)
		57967: 375, // VARIABLES (3x)
		57971: 376, // VIEW_SYM (3x)
		58000: 377, // VISIBLE_SYM (3x)
		57973: 378, // WAIT_SYM (3x)
		57974: 379, // WARNINGS (3x)
		57975: 380, // WEEK_SYM (3x)
		57976: 381, // WEIGHT_STRING_SYM (3x)
		57983: 382, // WITHOUT_SYM (3x)
		57984: 383, // WORK_SYM (3x)
		57985: 384, // WRAPPER_SYM (3x)
		57987: 385, // X509_SYM (3x)
		57988: 386, // XA_SYM (3x)
		57989: 387, // XID_SYM (3x)
		57990: 388, // XML_SYM (3x)
		57993: 389, // YEAR_SYM (3x)
		58009: 390, // $@1 (1x)
		58013: 391, // alter (1x)
		57357: 392, // ALTER (1x)
		58018: 393, // opt_end_of_input (1x)
		58021: 394, // simple_statement (1x)
		58022: 395, // simple_statement_or_begin (1x)
		58023: 396, // sql_statement (1x)
		58024: 397, // start_entry (1x)
		58025: 398, // table_ident (1x)
		57909: 399, // TABLE_SYM (1x)
		58010: 400, // $default (0x)
		37:    401, // '%' (0x)
		38:    402, // '&' (0x)
		40:    403, // '(' (0x)
		41:    404, // ')' (0x)
		42:    405, // '*' (0x)
		43:    406, // '+' (0x)
		45:    407, // '-' (0x)
		47:    408, // '/' (0x)
		94:    409, // '^' (0x)
		124:   410, // '|' (0x)
		126:   411, // '~' (0x)
		57346: 412, // ABORT_SYM (0x)
		57347: 413, // ACCESSIBLE_SYM (0x)
		57350: 414, // ADD (0x)
		57998: 415, // ADMIN_SYM (0x)
		57356: 416, // ALL (0x)
		57360: 417, // ANALYZE_SYM (0x)
		57361: 418, // AND_AND_SYM (0x)
		57362: 419, // AND_SYM (0x)
		57364: 420, // AS (0x)
		57365: 421, // ASC (0x)
		57367: 422, // ASENSITIVE_SYM (0x)
		57374: 423, // BEFORE_SYM (0x)
		57376: 424, // BETWEEN_SYM (0x)
		57377: 425, // BIGINT_SYM (0x)
		57380: 426, // BIN_NUM (0x)
		57378: 427, // BINARY_SYM (0x)
		57381: 428, // BIT_AND (0x)
		57382: 429, // BIT_OR (0x)
		57384: 430, // BIT_XOR (0x)
		57385: 431, // BLOB_SYM (0x)
		57389: 432, // BOTH (0x)
		57391: 433, // BY (0x)
		57394: 434, // CALL_SYM (0x)
		57395: 435, // CASCADE (0x)
		57397: 436, // CASE_SYM (0x)
		57398: 437, // CAST_SYM (0x)
		57401: 438, // CHANGE (0x)
		57405: 439, // CHAR_SYM (0x)
		57407: 440, // CHECK_SYM (0x)
		57414: 441, // COLLATE_SYM (0x)
		57417: 442, // COLUMN_SYM (0x)
		57429: 443, // CONDITION_SYM (0x)
		58006: 444, // CONDITIONLESS_JOIN (0x)
		57432: 445, // CONSTRAINT (0x)
		57438: 446, // CONTINUE_SYM (0x)
		57439: 447, // CONVERT_SYM (0x)
		57440: 448, // COUNT_SYM (0x)
		57442: 449, // CREATE (0x)
		57443: 450, // CROSS (0x)
		57445: 451, // CURDATE (0x)
		57447: 452, // CURRENT_USER (0x)
		57448: 453, // CURSOR_SYM (0x)
		57450: 454, // CURTIME (0x)
		57451: 455, // DATABASE (0x)
		57452: 456, // DATABASES (0x)
		57456: 457, // DATE_ADD_INTERVAL (0x)
		57457: 458, // DATE_SUB_INTERVAL (0x)
		57459: 459, // DAY_HOUR_SYM (0x)
		57460: 460, // DAY_MICROSECOND_SYM (0x)
		57461: 461, // DAY_MINUTE_SYM (0x)
		57462: 462, // DAY_SECOND_SYM (0x)
		57465: 463, // DECIMAL_NUM (0x)
		57466: 464, // DECIMAL_SYM (0x)
		57467: 465, // DECLARE_SYM (0x)
		57468: 466, // DEFAULT_SYM (0x)
		57471: 467, // DELAYED_SYM (0x)
		57473: 468, // DELETE_SYM (0x)
		57474: 469, // DESC (0x)
		57475: 470, // DESCRIBE (0x)
		57477: 471, // DETERMINISTIC_SYM (0x)
		57483: 472, // DISTINCT (0x)
		57484: 473, // DIV_SYM (0x)
		57485: 474, // DOUBLE_SYM (0x)
		57487: 475, // DROP (0x)
		57488: 476, // DUAL_SYM (0x)
		57492: 477, // EACH_SYM (0x)
		57493: 478, // ELSE (0x)
		57494: 479, // ELSEIF_SYM (0x)
		58008: 480, // EMPTY_FROM_CLAUSE (0x)
		57496: 481, // ENCLOSED (0x)
		57503: 482, // EQ (0x)
		57504: 483, // EQUAL_SYM (0x)
		57345: 484, // error (0x)
		57507: 485, // ESCAPED (0x)
		58001: 486, // EXCEPT_SYM (0x)
		57514: 487, // EXISTS (0x)
		57515: 488, // EXIT_SYM (0x)
		57521: 489, // EXTRACT_SYM (0x)
		57522: 490, // FALSE_SYM (0x)
		57525: 491, // FETCH_SYM (0x)
		57531: 492, // FLOAT_NUM (0x)
		57532: 493, // FLOAT_SYM (0x)
		57537: 494, // FOR_SYM (0x)
		57535: 495, // FORCE_SYM (0x)
		57536: 496, // FOREIGN (0x)
		57540: 497, // FROM (0x)
		57542: 498, // FULLTEXT_SYM (0x)
		57544: 499, // GE (0x)
		57546: 500, // GENERATED (0x)
		57551: 501, // GET_SYM (0x)
		58003: 502, // GRAMMAR_SELECTOR_EXPR (0x)
		58004: 503, // GRAMMAR_SELECTOR_GCOL (0x)
		58005: 504, // GRAMMAR_SELECTOR_PART (0x)
		57553: 505, // GRANT (0x)
		57556: 506, // GROUP_CONCAT_SYM (0x)
		57555: 507, // GROUP_SYM (0x)
		57557: 508, // GT_SYM (0x)
		57560: 509, // HAVING (0x)
		57562: 510, // HEX_NUM (0x)
		57563: 511, // HIGH_PRIORITY (0x)
		57566: 512, // HOUR_MICROSECOND_SYM (0x)
		57567: 513, // HOUR_MINUTE_SYM (0x)
		57568: 514, // HOUR_SECOND_SYM (0x)
		58016: 515, // ident_or_text (0x)
		57573: 516, // IF (0x)
		57574: 517, // IGNORE_SYM (0x)
		57592: 518, // IN_SYM (0x)
		57578: 519, // INDEX_SYM (0x)
		57579: 520, // INFILE (0x)
		57581: 521, // INNER_SYM (0x)
		57582: 522, // INOUT_SYM (0x)
		57583: 523, // INSENSITIVE_SYM (0x)
		57584: 524, // INSERT_SYM (0x)
		57590: 525, // INT_SYM (0x)
		57588: 526, // INTERVAL_SYM (0x)
		57589: 527, // INTO (0x)
		57593: 528, // IO_AFTER_GTIDS (0x)
		57594: 529, // IO_BEFORE_GTIDS (0x)
		57597: 530, // IS (0x)
		57600: 531, // ITERATE_SYM (0x)
		57601: 532, // JOIN_SYM (0x)
		57602: 533, // JSON_SEPARATOR_SYM (0x)
		57995: 534, // JSON_UNQUOTED_SEPARATOR_SYM (0x)
		57606: 535, // KEY_SYM (0x)
		57604: 536, // KEYS (0x)
		57607: 537, // KILL_SYM (0x)
		57610: 538, // LE (0x)
		57611: 539, // LEADING (0x)
		57613: 540, // LEAVE_SYM (0x)
		57614: 541, // LEFT (0x)
		57617: 542, // LEX_HOSTNAME (0x)
		57618: 543, // LIKE (0x)
		57619: 544, // LIMIT (0x)
		57620: 545, // LINEAR_SYM (0x)
		57621: 546, // LINES (0x)
		57624: 547, // LOAD (0x)
		57626: 548, // LOCATOR_SYM (0x)
		57628: 549, // LOCK_SYM (0x)
		57633: 550, // LONG_NUM (0x)
		57634: 551, // LONG_SYM (0x)
		57631: 552, // LONGBLOB_SYM (0x)
		57632: 553, // LONGTEXT_SYM (0x)
		57635: 554, // LOOP_SYM (0x)
		57636: 555, // LOW_PRIORITY (0x)
		57637: 556, // LT (0x)
		57639: 557, // MASTER_BIND_SYM (0x)
		57658: 558, // MASTER_SSL_VERIFY_SERVER_CERT_SYM (0x)
		57662: 559, // MATCH (0x)
		57667: 560, // MAX_SYM (0x)
		57670: 561, // MAX_VALUE_SYM (0x)
		57671: 562, // MEDIUMBLOB_SYM (0x)
		57672: 563, // MEDIUMINT_SYM (0x)
		57673: 564, // MEDIUMTEXT_SYM (0x)
		57684: 565, // MIN_SYM (0x)
		57680: 566, // MINUTE_MICROSECOND_SYM (0x)
		57681: 567, // MINUTE_SECOND_SYM (0x)
		57688: 568, // MOD_SYM (0x)
		57686: 569, // MODIFIES_SYM (0x)
		57698: 570, // NATURAL (0x)
		57699: 571, // NCHAR_STRING (0x)
		57702: 572, // NE (0x)
		57703: 573, // NEG (0x)
		57714: 574, // NO_WRITE_TO_BINLOG (0x)
		57709: 575, // NOT2_SYM (0x)
		57710: 576, // NOT_SYM (0x)
		57711: 577, // NOW_SYM (0x)
		57715: 578, // NULL_SYM (0x)
		57716: 579, // NUM (0x)
		57718: 580, // NUMERIC_SYM (0x)
		57742: 581, // OBSOLETE_TOKEN_654 (0x)
		57908: 582, // OBSOLETE_TOKEN_820 (0x)
		57721: 583, // ON_SYM (0x)
		57725: 584, // OPTIMIZE (0x)
		57726: 585, // OPTIMIZER_COSTS_SYM (0x)
		57728: 586, // OPTION (0x)
		57729: 587, // OPTIONALLY (0x)
		57730: 588, // OR2_SYM (0x)
		57732: 589, // OR_OR_SYM (0x)
		57733: 590, // OR_SYM (0x)
		57731: 591, // ORDER_SYM (0x)
		57736: 592, // OUT_SYM (0x)
		57734: 593, // OUTER (0x)
		57735: 594, // OUTFILE (0x)
		57740: 595, // PARAM_MARKER (0x)
		57744: 596, // PARTITION_SYM (0x)
		57996: 597, // PERSIST_SYM (0x)
		57755: 598, // POSITION_SYM (0x)
		57757: 599, // PRECISION (0x)
		57761: 600, // PRIMARY_SYM (0x)
		57763: 601, // PROCEDURE_SYM (0x)
		57769: 602, // PURGE (0x)
		57773: 603, // RANGE_SYM (0x)
		57776: 604, // READ_SYM (0x)
		57777: 605, // READ_WRITE_SYM (0x)
		57774: 606, // READS_SYM (0x)
		57778: 607, // REAL_SYM (0x)
		57784: 608, // REFERENCES (0x)
		57785: 609, // REGEXP (0x)
		57791: 610, // RELEASE_SYM (0x)
		57794: 611, // RENAME (0x)
		57798: 612, // REPEAT_SYM (0x)
		57799: 613, // REPLACE_SYM (0x)
		57808: 614, // REQUIRE_SYM (0x)
		57810: 615, // RESIGNAL_SYM (0x)
		57813: 616, // RESTRICT (0x)
		57817: 617, // RETURN_SYM (0x)
		57819: 618, // REVOKE (0x)
		57820: 619, // RIGHT (0x)
		57997: 620, // ROLE_SYM (0x)
		57833: 621, // SECOND_MICROSECOND_SYM (0x)
		57836: 622, // SELECT_SYM (0x)
		57837: 623, // SENSITIVE_SYM (0x)
		57838: 624, // SEPARATOR_SYM (0x)
		57843: 625, // SERVER_OPTIONS (0x)
		57844: 626, // SET_SYM (0x)
		57845: 627, // SET_VAR (0x)
		57847: 628, // SHIFT_LEFT (0x)
		57848: 629, // SHIFT_RIGHT (0x)
		57849: 630, // SHOW (0x)
		57851: 631, // SIGNAL_SYM (0x)
		57856: 632, // SMALLINT_SYM (0x)
		57862: 633, // SPATIAL_SYM (0x)
		57863: 634, // SPECIFIC_SYM (0x)
		57870: 635, // SQL_BIG_RESULT (0x)
		57873: 636, // SQL_CALC_FOUND_ROWS (0x)
		57875: 637, // SQL_SMALL_RESULT (0x)
		57876: 638, // SQL_SYM (0x)
		57864: 639, // SQLEXCEPTION_SYM (0x)
		57865: 640, // SQLSTATE_SYM (0x)
		57866: 641, // SQLWARNING_SYM (0x)
		57878: 642, // SSL_SYM (0x)
		57880: 643, // STARTING (0x)
		57888: 644, // STD_SYM (0x)
		57887: 645, // STDDEV_SAMP_SYM (0x)
		57891: 646, // STORED_SYM (0x)
		57892: 647, // STRAIGHT_JOIN (0x)
		58007: 648, // SUBQUERY_AS_EXPR (0x)
		57899: 649, // SUBSTRING (0x)
		57900: 650, // SUM_SYM (0x)
		57905: 651, // SYSDATE (0x)
		57914: 652, // TERMINATED (0x)
		57915: 653, // TEXT_STRING (0x)
		58012: 654, // TEXT_STRING_sys (0x)
		57918: 655, // THEN_SYM (0x)
		57923: 656, // TINYBLOB_SYM (0x)
		57924: 657, // TINYINT_SYM (0x)
		57925: 658, // TINYTEXT_SYN (0x)
		57926: 659, // TO_SYM (0x)
		57927: 660, // TRAILING (0x)
		57930: 661, // TRIGGER_SYM (0x)
		57931: 662, // TRIM (0x)
		57932: 663, // TRUE_SYM (0x)
		57937: 664, // ULONGLONG_NUM (0x)
		57940: 665, // UNDERSCORE_CHARSET (0x)
		57943: 666, // UNDO_SYM (0x)
		57946: 667, // UNION_SYM (0x)
		57947: 668, // UNIQUE_SYM (0x)
		57949: 669, // UNLOCK_SYM (0x)
		57950: 670, // UNSIGNED_SYM (0x)
		57952: 671, // UPDATE_SYM (0x)
		57954: 672, // USAGE (0x)
		57957: 673, // USE_SYM (0x)
		57958: 674, // USING (0x)
		57959: 675, // UTC_DATE_SYM (0x)
		57961: 676, // UTC_TIME_SYM (0x)
		57960: 677, // UTC_TIMESTAMP_SYM (0x)
		57963: 678, // VALUES (0x)
		57970: 679, // VAR_SAMP_SYM (0x)
		57965: 680, // VARBINARY_SYM (0x)
		57966: 681, // VARCHAR_SYM (0x)
		57968: 682, // VARIANCE_SYM (0x)
		57969: 683, // VARYING (0x)
		57972: 684, // VIRTUAL_SYM (0x)
		57977: 685, // WHEN_SYM (0x)
		57978: 686, // WHERE (0x)
		57979: 687, // WHILE_SYM (0x)
		57980: 688, // WITH (0x)
		57981: 689, // WITH_CUBE_SYM (0x)
		57982: 690, // WITH_ROLLUP_SYM (0x)
		57986: 691, // WRITE_SYM (0x)
		57991: 692, // XOR (0x)
		57992: 693, // YEAR_MONTH_SYM (0x)
		57994: 694, // ZEROFILL_SYM (0x)
	}

	MySQLSymNames = []string{
		"END_OF_INPUT",
		"';'",
		"'.'",
		"$end",
		"ACCOUNT_SYM",
		"ACTION",
		"ADDDATE_SYM",
		"AFTER_SYM",
		"AGAINST",
		"AGGREGATE_SYM",
		"ALGORITHM_SYM",
		"ALWAYS_SYM",
		"ANALYSE_SYM",
		"ANY_SYM",
		"ASCII_SYM",
		"AT_SYM",
		"AUTO_INC",
		"AUTOEXTEND_SIZE_SYM",
		"AVG_ROW_LENGTH",
		"AVG_SYM",
		"BACKUP_SYM",
		"BEGIN_SYM",
		"BINLOG_SYM",
		"BIT_SYM",
		"BLOCK_SYM",
		"BOOL_SYM",
		"BOOLEAN_SYM",
		"BTREE_SYM",
		"BYTE_SYM",
		"CACHE_SYM",
		"CASCADED",
		"CATALOG_NAME_SYM",
		"CHAIN_SYM",
		"CHANGED",
		"CHANNEL_SYM",
		"CHARSET",
		"CHECKSUM_SYM",
		"CIPHER_SYM",
		"CLASS_ORIGIN_SYM",
		"CLIENT_SYM",
		"CLOSE_SYM",
		"COALESCE",
		"CODE_SYM",
		"COLLATION_SYM",
		"COLUMN_FORMAT_SYM",
		"COLUMN_NAME_SYM",
		"COLUMNS",
		"COMMENT_SYM",
		"COMMIT_SYM",
		"COMMITTED_SYM",
		"COMPACT_SYM",
		"COMPLETION_SYM",
		"COMPONENT_SYM",
		"COMPRESSED_SYM",
		"COMPRESSION_SYM",
		"CONCURRENT",
		"CONNECTION_SYM",
		"CONSISTENT_SYM",
		"CONSTRAINT_CATALOG_SYM",
		"CONSTRAINT_NAME_SYM",
		"CONSTRAINT_SCHEMA_SYM",
		"CONTAINS_SYM",
		"CONTEXT_SYM",
		"CPU_SYM",
		"CUBE_SYM",
		"CURRENT_SYM",
		"CURSOR_NAME_SYM",
		"DATA_SYM",
		"DATAFILE_SYM",
		"DATE_SYM",
		"DATETIME_SYM",
		"DAY_SYM",
		"DEALLOCATE_SYM",
		"DEFAULT_AUTH_SYM",
		"DEFINER_SYM",
		"DELAY_KEY_WRITE_SYM",
		"DES_KEY_FILE",
		"DIAGNOSTICS_SYM",
		"DIRECTORY_SYM",
		"DISABLE_SYM",
		"DISCARD",
		"DISK_SYM",
		"DO_SYM",
		"DUMPFILE",
		"DUPLICATE_SYM",
		"DYNAMIC_SYM",
		"ENABLE_SYM",
		"ENCRYPTION_SYM",
		"END",
		"ENDS_SYM",
		"ENGINE_SYM",
		"ENGINES_SYM",
		"ENUM_SYM",
		"ERROR_SYM",
		"ERRORS",
		"ESCAPE_SYM",
		"EVENT_SYM",
		"EVENTS_SYM",
		"EVERY_SYM",
		"EXCHANGE_SYM",
		"EXECUTE_SYM",
		"EXPANSION_SYM",
		"EXPIRE_SYM",
		"EXPORT_SYM",
		"EXTENDED_SYM",
		"EXTENT_SIZE_SYM",
		"FAST_SYM",
		"FAULTS_SYM",
		"FILE_BLOCK_SIZE_SYM",
		"FILE_SYM",
		"FILTER_SYM",
		"FIRST_SYM",
		"FIXED_SYM",
		"FLUSH_SYM",
		"FOLLOWS_SYM",
		"FORMAT_SYM",
		"FOUND_SYM",
		"FULL",
		"FUNCTION_SYM",
		"GENERAL",
		"GEOMETRY_SYM",
		"GEOMETRYCOLLECTION_SYM",
		"GET_FORMAT",
		"GLOBAL_SYM",
		"GRANTS",
		"GROUP_REPLICATION",
		"HANDLER_SYM",
		"HASH_SYM",
		"HELP_SYM",
		"HOST_SYM",
		"HOSTS_SYM",
		"HOUR_SYM",
		"ident",
		"IDENT",
		"ident_keyword",
		"IDENT_QUOTED",
		"IDENT_sys",
		"IDENTIFIED_SYM",
		"IGNORE_SERVER_IDS_SYM",
		"IMPORT",
		"INDEXES",
		"INITIAL_SIZE_SYM",
		"INSERT_METHOD",
		"INSTALL_SYM",
		"INSTANCE_SYM",
		"INVISIBLE_SYM",
		"INVOKER_SYM",
		"IO_SYM",
		"IPC_SYM",
		"ISOLATION",
		"ISSUER_SYM",
		"JSON_SYM",
		"KEY_BLOCK_SIZE",
		"label_keyword",
		"LANGUAGE_SYM",
		"LAST_SYM",
		"LEAVES",
		"LESS_SYM",
		"LEVEL_SYM",
		"LINESTRING_SYM",
		"LIST_SYM",
		"LOCAL_SYM",
		"LOCKS_SYM",
		"LOGFILE_SYM",
		"LOGS_SYM",
		"MASTER_AUTO_POSITION_SYM",
		"MASTER_CONNECT_RETRY_SYM",
		"MASTER_DELAY_SYM",
		"MASTER_HEARTBEAT_PERIOD_SYM",
		"MASTER_HOST_SYM",
		"MASTER_LOG_FILE_SYM",
		"MASTER_LOG_POS_SYM",
		"MASTER_PASSWORD_SYM",
		"MASTER_PORT_SYM",
		"MASTER_RETRY_COUNT_SYM",
		"MASTER_SERVER_ID_SYM",
		"MASTER_SSL_CA_SYM",
		"MASTER_SSL_CAPATH_SYM",
		"MASTER_SSL_CERT_SYM",
		"MASTER_SSL_CIPHER_SYM",
		"MASTER_SSL_CRL_SYM",
		"MASTER_SSL_CRLPATH_SYM",
		"MASTER_SSL_KEY_SYM",
		"MASTER_SSL_SYM",
		"MASTER_SYM",
		"MASTER_TLS_VERSION_SYM",
		"MASTER_USER_SYM",
		"MAX_CONNECTIONS_PER_HOUR",
		"MAX_QUERIES_PER_HOUR",
		"MAX_ROWS",
		"MAX_SIZE_SYM",
		"MAX_UPDATES_PER_HOUR",
		"MAX_USER_CONNECTIONS_SYM",
		"MEDIUM_SYM",
		"MEMORY_SYM",
		"MERGE_SYM",
		"MESSAGE_TEXT_SYM",
		"MICROSECOND_SYM",
		"MIGRATE_SYM",
		"MIN_ROWS",
		"MINUTE_SYM",
		"MODE_SYM",
		"MODIFY_SYM",
		"MONTH_SYM",
		"MULTILINESTRING_SYM",
		"MULTIPOINT_SYM",
		"MULTIPOLYGON_SYM",
		"MUTEX_SYM",
		"MYSQL_ERRNO_SYM",
		"NAME_SYM",
		"NAMES_SYM",
		"NATIONAL_SYM",
		"NCHAR_SYM",
		"NDBCLUSTER_SYM",
		"NEVER_SYM",
		"NEW_SYM",
		"NEXT_SYM",
		"NO_SYM",
		"NO_WAIT_SYM",
		"NODEGROUP_SYM",
		"NONE_SYM",
		"NUMBER_SYM",
		"NVARCHAR_SYM",
		"OFFSET_SYM",
		"ONE_SYM",
		"ONLY_SYM",
		"OPEN_SYM",
		"OPTIONS_SYM",
		"OWNER_SYM",
		"PACK_KEYS_SYM",
		"PAGE_SYM",
		"PARSER_SYM",
		"PARTIAL",
		"PARTITIONING_SYM",
		"PARTITIONS_SYM",
		"PASSWORD",
		"PHASE_SYM",
		"PLUGIN_DIR_SYM",
		"PLUGIN_SYM",
		"PLUGINS_SYM",
		"POINT_SYM",
		"POLYGON_SYM",
		"PORT_SYM",
		"PRECEDES_SYM",
		"PREPARE_SYM",
		"PRESERVE_SYM",
		"PREV_SYM",
		"PRIVILEGES",
		"PROCESS",
		"PROCESSLIST_SYM",
		"PROFILE_SYM",
		"PROFILES_SYM",
		"PROXY_SYM",
		"QUARTER_SYM",
		"QUERY_SYM",
		"QUICK",
		"READ_ONLY_SYM",
		"REBUILD_SYM",
		"RECOVER_SYM",
		"REDO_BUFFER_SIZE_SYM",
		"REDOFILE_SYM",
		"REDUNDANT_SYM",
		"RELAY",
		"RELAY_LOG_FILE_SYM",
		"RELAY_LOG_POS_SYM",
		"RELAY_THREAD",
		"RELAYLOG_SYM",
		"RELOAD",
		"REMOVE_SYM",
		"REORGANIZE_SYM",
		"REPAIR",
		"REPEATABLE_SYM",
		"REPLICATE_DO_DB",
		"REPLICATE_DO_TABLE",
		"REPLICATE_IGNORE_DB",
		"REPLICATE_IGNORE_TABLE",
		"REPLICATE_REWRITE_DB",
		"REPLICATE_WILD_DO_TABLE",
		"REPLICATE_WILD_IGNORE_TABLE",
		"REPLICATION",
		"RESET_SYM",
		"RESOURCES",
		"RESTORE_SYM",
		"RESUME_SYM",
		"RETURNED_SQLSTATE_SYM",
		"RETURNS_SYM",
		"REVERSE_SYM",
		"role_or_ident_keyword",
		"role_or_label_keyword",
		"ROLLBACK_SYM",
		"ROLLUP_SYM",
		"ROTATE_SYM",
		"ROUTINE_SYM",
		"ROW_COUNT_SYM",
		"ROW_FORMAT_SYM",
		"ROW_SYM",
		"ROWS_SYM",
		"RTREE_SYM",
		"SAVEPOINT_SYM",
		"SCHEDULE_SYM",
		"SCHEMA_NAME_SYM",
		"SECOND_SYM",
		"SECURITY_SYM",
		"SERIAL_SYM",
		"SERIALIZABLE_SYM",
		"SERVER_SYM",
		"SESSION_SYM",
		"SHARE_SYM",
		"SHUTDOWN",
		"SIGNED_SYM",
		"SIMPLE_SYM",
		"SLAVE",
		"SLOW",
		"SNAPSHOT_SYM",
		"SOCKET_SYM",
		"SONAME_SYM",
		"SOUNDS_SYM",
		"SOURCE_SYM",
		"SQL_AFTER_GTIDS",
		"SQL_AFTER_MTS_GAPS",
		"SQL_BEFORE_GTIDS",
		"SQL_BUFFER_RESULT",
		"SQL_CACHE_SYM",
		"SQL_NO_CACHE_SYM",
		"SQL_THREAD",
		"STACKED_SYM",
		"START_SYM",
		"STARTS_SYM",
		"STATS_AUTO_RECALC_SYM",
		"STATS_PERSISTENT_SYM",
		"STATS_SAMPLE_PAGES_SYM",
		"STATUS_SYM",
		"STOP_SYM",
		"STORAGE_SYM",
		"STRING_SYM",
		"SUBCLASS_ORIGIN_SYM",
		"SUBDATE_SYM",
		"SUBJECT_SYM",
		"SUBPARTITION_SYM",
		"SUBPARTITIONS_SYM",
		"SUPER_SYM",
		"SUSPEND_SYM",
		"SWAPS_SYM",
		"SWITCHES_SYM",
		"TABLE_CHECKSUM_SYM",
		"TABLE_NAME_SYM",
		"TABLES",
		"TABLESPACE_SYM",
		"TEMPORARY",
		"TEMPTABLE_SYM",
		"TEXT_SYM",
		"THAN_SYM",
		"TIME_SYM",
		"TIMESTAMP_ADD",
		"TIMESTAMP_DIFF",
		"TIMESTAMP_SYM",
		"TRANSACTION_SYM",
		"TRIGGERS_SYM",
		"TRUNCATE_SYM",
		"TYPE_SYM",
		"TYPES_SYM",
		"UDF_RETURNS_SYM",
		"UNCOMMITTED_SYM",
		"UNDEFINED_SYM",
		"UNDO_BUFFER_SIZE_SYM",
		"UNDOFILE_SYM",
		"UNICODE_SYM",
		"UNINSTALL_SYM",
		"UNKNOWN_SYM",
		"UNTIL_SYM",
		"UPGRADE_SYM",
		"USE_FRM",
		"USER",
		"VALIDATION_SYM",
		"VALUE_SYM",
		"VARIABLES",
		"VIEW_SYM",
		"VISIBLE_SYM",
		"WAIT_SYM",
		"WARNINGS",
		"WEEK_SYM",
		"WEIGHT_STRING_SYM",
		"WITHOUT_SYM",
		"WORK_SYM",
		"WRAPPER_SYM",
		"X509_SYM",
		"XA_SYM",
		"XID_SYM",
		"XML_SYM",
		"YEAR_SYM",
		"$@1",
		"alter",
		"ALTER",
		"opt_end_of_input",
		"simple_statement",
		"simple_statement_or_begin",
		"sql_statement",
		"start_entry",
		"table_ident",
		"TABLE_SYM",
		"$default",
		"'%'",
		"'&'",
		"'('",
		"')'",
		"'*'",
		"'+'",
		"'-'",
		"'/'",
		"'^'",
		"'|'",
		"'~'",
		"ABORT_SYM",
		"ACCESSIBLE_SYM",
		"ADD",
		"ADMIN_SYM",
		"ALL",
		"ANALYZE_SYM",
		"AND_AND_SYM",
		"AND_SYM",
		"AS",
		"ASC",
		"ASENSITIVE_SYM",
		"BEFORE_SYM",
		"BETWEEN_SYM",
		"BIGINT_SYM",
		"BIN_NUM",
		"BINARY_SYM",
		"BIT_AND",
		"BIT_OR",
		"BIT_XOR",
		"BLOB_SYM",
		"BOTH",
		"BY",
		"CALL_SYM",
		"CASCADE",
		"CASE_SYM",
		"CAST_SYM",
		"CHANGE",
		"CHAR_SYM",
		"CHECK_SYM",
		"COLLATE_SYM",
		"COLUMN_SYM",
		"CONDITION_SYM",
		"CONDITIONLESS_JOIN",
		"CONSTRAINT",
		"CONTINUE_SYM",
		"CONVERT_SYM",
		"COUNT_SYM",
		"CREATE",
		"CROSS",
		"CURDATE",
		"CURRENT_USER",
		"CURSOR_SYM",
		"CURTIME",
		"DATABASE",
		"DATABASES",
		"DATE_ADD_INTERVAL",
		"DATE_SUB_INTERVAL",
		"DAY_HOUR_SYM",
		"DAY_MICROSECOND_SYM",
		"DAY_MINUTE_SYM",
		"DAY_SECOND_SYM",
		"DECIMAL_NUM",
		"DECIMAL_SYM",
		"DECLARE_SYM",
		"DEFAULT_SYM",
		"DELAYED_SYM",
		"DELETE_SYM",
		"DESC",
		"DESCRIBE",
		"DETERMINISTIC_SYM",
		"DISTINCT",
		"DIV_SYM",
		"DOUBLE_SYM",
		"DROP",
		"DUAL_SYM",
		"EACH_SYM",
		"ELSE",
		"ELSEIF_SYM",
		"EMPTY_FROM_CLAUSE",
		"ENCLOSED",
		"EQ",
		"EQUAL_SYM",
		"error",
		"ESCAPED",
		"EXCEPT_SYM",
		"EXISTS",
		"EXIT_SYM",
		"EXTRACT_SYM",
		"FALSE_SYM",
		"FETCH_SYM",
		"FLOAT_NUM",
		"FLOAT_SYM",
		"FOR_SYM",
		"FORCE_SYM",
		"FOREIGN",
		"FROM",
		"FULLTEXT_SYM",
		"GE",
		"GENERATED",
		"GET_SYM",
		"GRAMMAR_SELECTOR_EXPR",
		"GRAMMAR_SELECTOR_GCOL",
		"GRAMMAR_SELECTOR_PART",
		"GRANT",
		"GROUP_CONCAT_SYM",
		"GROUP_SYM",
		"GT_SYM",
		"HAVING",
		"HEX_NUM",
		"HIGH_PRIORITY",
		"HOUR_MICROSECOND_SYM",
		"HOUR_MINUTE_SYM",
		"HOUR_SECOND_SYM",
		"ident_or_text",
		"IF",
		"IGNORE_SYM",
		"IN_SYM",
		"INDEX_SYM",
		"INFILE",
		"INNER_SYM",
		"INOUT_SYM",
		"INSENSITIVE_SYM",
		"INSERT_SYM",
		"INT_SYM",
		"INTERVAL_SYM",
		"INTO",
		"IO_AFTER_GTIDS",
		"IO_BEFORE_GTIDS",
		"IS",
		"ITERATE_SYM",
		"JOIN_SYM",
		"JSON_SEPARATOR_SYM",
		"JSON_UNQUOTED_SEPARATOR_SYM",
		"KEY_SYM",
		"KEYS",
		"KILL_SYM",
		"LE",
		"LEADING",
		"LEAVE_SYM",
		"LEFT",
		"LEX_HOSTNAME",
		"LIKE",
		"LIMIT",
		"LINEAR_SYM",
		"LINES",
		"LOAD",
		"LOCATOR_SYM",
		"LOCK_SYM",
		"LONG_NUM",
		"LONG_SYM",
		"LONGBLOB_SYM",
		"LONGTEXT_SYM",
		"LOOP_SYM",
		"LOW_PRIORITY",
		"LT",
		"MASTER_BIND_SYM",
		"MASTER_SSL_VERIFY_SERVER_CERT_SYM",
		"MATCH",
		"MAX_SYM",
		"MAX_VALUE_SYM",
		"MEDIUMBLOB_SYM",
		"MEDIUMINT_SYM",
		"MEDIUMTEXT_SYM",
		"MIN_SYM",
		"MINUTE_MICROSECOND_SYM",
		"MINUTE_SECOND_SYM",
		"MOD_SYM",
		"MODIFIES_SYM",
		"NATURAL",
		"NCHAR_STRING",
		"NE",
		"NEG",
		"NO_WRITE_TO_BINLOG",
		"NOT2_SYM",
		"NOT_SYM",
		"NOW_SYM",
		"NULL_SYM",
		"NUM",
		"NUMERIC_SYM",
		"OBSOLETE_TOKEN_654",
		"OBSOLETE_TOKEN_820",
		"ON_SYM",
		"OPTIMIZE",
		"OPTIMIZER_COSTS_SYM",
		"OPTION",
		"OPTIONALLY",
		"OR2_SYM",
		"OR_OR_SYM",
		"OR_SYM",
		"ORDER_SYM",
		"OUT_SYM",
		"OUTER",
		"OUTFILE",
		"PARAM_MARKER",
		"PARTITION_SYM",
		"PERSIST_SYM",
		"POSITION_SYM",
		"PRECISION",
		"PRIMARY_SYM",
		"PROCEDURE_SYM",
		"PURGE",
		"RANGE_SYM",
		"READ_SYM",
		"READ_WRITE_SYM",
		"READS_SYM",
		"REAL_SYM",
		"REFERENCES",
		"REGEXP",
		"RELEASE_SYM",
		"RENAME",
		"REPEAT_SYM",
		"REPLACE_SYM",
		"REQUIRE_SYM",
		"RESIGNAL_SYM",
		"RESTRICT",
		"RETURN_SYM",
		"REVOKE",
		"RIGHT",
		"ROLE_SYM",
		"SECOND_MICROSECOND_SYM",
		"SELECT_SYM",
		"SENSITIVE_SYM",
		"SEPARATOR_SYM",
		"SERVER_OPTIONS",
		"SET_SYM",
		"SET_VAR",
		"SHIFT_LEFT",
		"SHIFT_RIGHT",
		"SHOW",
		"SIGNAL_SYM",
		"SMALLINT_SYM",
		"SPATIAL_SYM",
		"SPECIFIC_SYM",
		"SQL_BIG_RESULT",
		"SQL_CALC_FOUND_ROWS",
		"SQL_SMALL_RESULT",
		"SQL_SYM",
		"SQLEXCEPTION_SYM",
		"SQLSTATE_SYM",
		"SQLWARNING_SYM",
		"SSL_SYM",
		"STARTING",
		"STD_SYM",
		"STDDEV_SAMP_SYM",
		"STORED_SYM",
		"STRAIGHT_JOIN",
		"SUBQUERY_AS_EXPR",
		"SUBSTRING",
		"SUM_SYM",
		"SYSDATE",
		"TERMINATED",
		"TEXT_STRING",
		"TEXT_STRING_sys",
		"THEN_SYM",
		"TINYBLOB_SYM",
		"TINYINT_SYM",
		"TINYTEXT_SYN",
		"TO_SYM",
		"TRAILING",
		"TRIGGER_SYM",
		"TRIM",
		"TRUE_SYM",
		"ULONGLONG_NUM",
		"UNDERSCORE_CHARSET",
		"UNDO_SYM",
		"UNION_SYM",
		"UNIQUE_SYM",
		"UNLOCK_SYM",
		"UNSIGNED_SYM",
		"UPDATE_SYM",
		"USAGE",
		"USE_SYM",
		"USING",
		"UTC_DATE_SYM",
		"UTC_TIME_SYM",
		"UTC_TIMESTAMP_SYM",
		"VALUES",
		"VAR_SAMP_SYM",
		"VARBINARY_SYM",
		"VARCHAR_SYM",
		"VARIANCE_SYM",
		"VARYING",
		"VIRTUAL_SYM",
		"WHEN_SYM",
		"WHERE",
		"WHILE_SYM",
		"WITH",
		"WITH_CUBE_SYM",
		"WITH_ROLLUP_SYM",
		"WRITE_SYM",
		"XOR",
		"YEAR_MONTH_SYM",
		"ZEROFILL_SYM",
	}

	MySQLTokenLiteralStrings = map[int]string{}

	MySQLReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {397, 1},
		2:   {396, 1},
		3:   {390, 0},
		4:   {396, 4},
		5:   {396, 2},
		6:   {393, 0},
		7:   {393, 1},
		8:   {395, 1},
		9:   {394, 1},
		10:  {391, 3},
		11:  {398, 1},
		12:  {398, 3},
		13:  {398, 2},
		14:  {132, 1},
		15:  {132, 1},
		16:  {136, 1},
		17:  {136, 1},
		18:  {134, 1},
		19:  {134, 1},
		20:  {134, 1},
		21:  {134, 1},
		22:  {153, 1},
		23:  {153, 1},
		24:  {153, 1},
		25:  {153, 1},
		26:  {153, 1},
		27:  {153, 1},
		28:  {153, 1},
		29:  {153, 1},
		30:  {153, 1},
		31:  {287, 1},
		32:  {287, 1},
		33:  {287, 1},
		34:  {287, 1},
		35:  {287, 1},
		36:  {287, 1},
		37:  {287, 1},
		38:  {287, 1},
		39:  {287, 1},
		40:  {287, 1},
		41:  {287, 1},
		42:  {287, 1},
		43:  {287, 1},
		44:  {287, 1},
		45:  {287, 1},
		46:  {287, 1},
		47:  {287, 1},
		48:  {287, 1},
		49:  {287, 1},
		50:  {287, 1},
		51:  {287, 1},
		52:  {287, 1},
		53:  {287, 1},
		54:  {287, 1},
		55:  {287, 1},
		56:  {287, 1},
		57:  {287, 1},
		58:  {287, 1},
		59:  {287, 1},
		60:  {287, 1},
		61:  {287, 1},
		62:  {287, 1},
		63:  {287, 1},
		64:  {287, 1},
		65:  {287, 1},
		66:  {287, 1},
		67:  {287, 1},
		68:  {287, 1},
		69:  {287, 1},
		70:  {287, 1},
		71:  {287, 1},
		72:  {287, 1},
		73:  {287, 1},
		74:  {287, 1},
		75:  {287, 1},
		76:  {287, 1},
		77:  {287, 1},
		78:  {287, 1},
		79:  {287, 1},
		80:  {287, 1},
		81:  {287, 1},
		82:  {287, 1},
		83:  {287, 1},
		84:  {287, 1},
		85:  {287, 1},
		86:  {288, 1},
		87:  {288, 1},
		88:  {288, 1},
		89:  {288, 1},
		90:  {288, 1},
		91:  {288, 1},
		92:  {288, 1},
		93:  {288, 1},
		94:  {288, 1},
		95:  {288, 1},
		96:  {288, 1},
		97:  {288, 1},
		98:  {288, 1},
		99:  {288, 1},
		100: {288, 1},
		101: {288, 1},
		102: {288, 1},
		103: {288, 1},
		104: {288, 1},
		105: {288, 1},
		106: {288, 1},
		107: {288, 1},
		108: {288, 1},
		109: {288, 1},
		110: {288, 1},
		111: {288, 1},
		112: {288, 1},
		113: {288, 1},
		114: {288, 1},
		115: {288, 1},
		116: {288, 1},
		117: {288, 1},
		118: {288, 1},
		119: {288, 1},
		120: {288, 1},
		121: {288, 1},
		122: {288, 1},
		123: {288, 1},
		124: {288, 1},
		125: {288, 1},
		126: {288, 1},
		127: {288, 1},
		128: {288, 1},
		129: {288, 1},
		130: {288, 1},
		131: {288, 1},
		132: {288, 1},
		133: {288, 1},
		134: {288, 1},
		135: {288, 1},
		136: {288, 1},
		137: {288, 1},
		138: {288, 1},
		139: {288, 1},
		140: {288, 1},
		141: {288, 1},
		142: {288, 1},
		143: {288, 1},
		144: {288, 1},
		145: {288, 1},
		146: {288, 1},
		147: {288, 1},
		148: {288, 1},
		149: {288, 1},
		150: {288, 1},
		151: {288, 1},
		152: {288, 1},
		153: {288, 1},
		154: {288, 1},
		155: {288, 1},
		156: {288, 1},
		157: {288, 1},
		158: {288, 1},
		159: {288, 1},
		160: {288, 1},
		161: {288, 1},
		162: {288, 1},
		163: {288, 1},
		164: {288, 1},
		165: {288, 1},
		166: {288, 1},
		167: {288, 1},
		168: {288, 1},
		169: {288, 1},
		170: {288, 1},
		171: {288, 1},
		172: {288, 1},
		173: {288, 1},
		174: {288, 1},
		175: {288, 1},
		176: {288, 1},
		177: {288, 1},
		178: {288, 1},
		179: {288, 1},
		180: {288, 1},
		181: {288, 1},
		182: {288, 1},
		183: {288, 1},
		184: {288, 1},
		185: {288, 1},
		186: {288, 1},
		187: {288, 1},
		188: {288, 1},
		189: {288, 1},
		190: {288, 1},
		191: {288, 1},
		192: {288, 1},
		193: {288, 1},
		194: {288, 1},
		195: {288, 1},
		196: {288, 1},
		197: {288, 1},
		198: {288, 1},
		199: {288, 1},
		200: {288, 1},
		201: {288, 1},
		202: {288, 1},
		203: {288, 1},
		204: {288, 1},
		205: {288, 1},
		206: {288, 1},
		207: {288, 1},
		208: {288, 1},
		209: {288, 1},
		210: {288, 1},
		211: {288, 1},
		212: {288, 1},
		213: {288, 1},
		214: {288, 1},
		215: {288, 1},
		216: {288, 1},
		217: {288, 1},
		218: {288, 1},
		219: {288, 1},
		220: {288, 1},
		221: {288, 1},
		222: {288, 1},
		223: {288, 1},
		224: {288, 1},
		225: {288, 1},
		226: {288, 1},
		227: {288, 1},
		228: {288, 1},
		229: {288, 1},
		230: {288, 1},
		231: {288, 1},
		232: {288, 1},
		233: {288, 1},
		234: {288, 1},
		235: {288, 1},
		236: {288, 1},
		237: {288, 1},
		238: {288, 1},
		239: {288, 1},
		240: {288, 1},
		241: {288, 1},
		242: {288, 1},
		243: {288, 1},
		244: {288, 1},
		245: {288, 1},
		246: {288, 1},
		247: {288, 1},
		248: {288, 1},
		249: {288, 1},
		250: {288, 1},
		251: {288, 1},
		252: {288, 1},
		253: {288, 1},
		254: {288, 1},
		255: {288, 1},
		256: {288, 1},
		257: {288, 1},
		258: {288, 1},
		259: {288, 1},
		260: {288, 1},
		261: {288, 1},
		262: {288, 1},
		263: {288, 1},
		264: {288, 1},
		265: {288, 1},
		266: {288, 1},
		267: {288, 1},
		268: {288, 1},
		269: {288, 1},
		270: {288, 1},
		271: {288, 1},
		272: {288, 1},
		273: {288, 1},
		274: {288, 1},
		275: {288, 1},
		276: {288, 1},
		277: {288, 1},
		278: {288, 1},
		279: {288, 1},
		280: {288, 1},
		281: {288, 1},
		282: {288, 1},
		283: {288, 1},
		284: {288, 1},
		285: {288, 1},
		286: {288, 1},
		287: {288, 1},
		288: {288, 1},
		289: {288, 1},
		290: {288, 1},
		291: {288, 1},
		292: {288, 1},
		293: {288, 1},
		294: {288, 1},
		295: {288, 1},
		296: {288, 1},
		297: {288, 1},
		298: {288, 1},
		299: {288, 1},
		300: {288, 1},
		301: {288, 1},
		302: {288, 1},
		303: {288, 1},
		304: {288, 1},
		305: {288, 1},
		306: {288, 1},
		307: {288, 1},
		308: {288, 1},
		309: {288, 1},
		310: {288, 1},
		311: {288, 1},
		312: {288, 1},
		313: {288, 1},
		314: {288, 1},
		315: {288, 1},
		316: {288, 1},
		317: {288, 1},
		318: {288, 1},
		319: {288, 1},
		320: {288, 1},
		321: {288, 1},
		322: {288, 1},
		323: {288, 1},
		324: {288, 1},
		325: {288, 1},
		326: {288, 1},
		327: {288, 1},
		328: {288, 1},
		329: {288, 1},
		330: {288, 1},
		331: {288, 1},
		332: {288, 1},
		333: {288, 1},
		334: {288, 1},
		335: {288, 1},
		336: {288, 1},
		337: {288, 1},
		338: {288, 1},
		339: {288, 1},
		340: {288, 1},
		341: {288, 1},
		342: {288, 1},
		343: {288, 1},
		344: {288, 1},
		345: {288, 1},
		346: {288, 1},
		347: {288, 1},
		348: {288, 1},
		349: {288, 1},
		350: {288, 1},
		351: {288, 1},
		352: {288, 1},
		353: {288, 1},
		354: {288, 1},
		355: {288, 1},
		356: {288, 1},
		357: {288, 1},
		358: {288, 1},
		359: {288, 1},
		360: {288, 1},
		361: {288, 1},
		362: {288, 1},
		363: {288, 1},
		364: {288, 1},
		365: {288, 1},
		366: {288, 1},
		367: {288, 1},
		368: {288, 1},
		369: {288, 1},
		370: {288, 1},
		371: {288, 1},
		372: {288, 1},
		373: {288, 1},
		374: {288, 1},
		375: {288, 1},
		376: {288, 1},
		377: {288, 1},
		378: {288, 1},
		379: {288, 1},
		380: {288, 1},
		381: {288, 1},
		382: {288, 1},
		383: {288, 1},
		384: {288, 1},
		385: {288, 1},
		386: {288, 1},
		387: {288, 1},
		388: {288, 1},
		389: {288, 1},
		390: {288, 1},
		391: {288, 1},
		392: {288, 1},
		393: {288, 1},
		394: {288, 1},
		395: {288, 1},
		396: {288, 1},
		397: {288, 1},
		398: {288, 1},
		399: {515, 1},
		400: {515, 1},
		401: {515, 1},
		402: {654, 1},
	}

	MySQLXErrors = map[MySQLXError]string{}

	MySQLParseTab = [405][]uint16{
		// 0
		{402, 391: 405, 406, 394: 404, 403, 401, 400},
		{3: 399},
		{3: 398},
		{3: 397},
		{800, 396, 390: 799},
		// 5
		{391, 391},
		{390, 390},
		{399: 407},
		{2: 410, 4: 428, 483, 484, 485, 486, 487, 488, 430, 489, 490, 429, 491, 492, 493, 494, 495, 431, 432, 496, 497, 498, 499, 500, 501, 433, 434, 502, 503, 504, 505, 506, 435, 436, 507, 509, 508, 437, 510, 511, 512, 514, 513, 515, 438, 439, 516, 517, 518, 519, 520, 521, 523, 524, 525, 526, 528, 527, 440, 529, 530, 531, 532, 533, 534, 535, 537, 536, 538, 441, 539, 540, 541, 542, 543, 544, 545, 546, 547, 442, 548, 549, 550, 569, 522, 443, 551, 553, 554, 552, 555, 556, 557, 420, 558, 559, 560, 417, 561, 562, 563, 564, 565, 567, 566, 571, 421, 572, 573, 574, 444, 445, 446, 568, 570, 773, 575, 576, 577, 578, 580, 579, 447, 448, 581, 449, 450, 582, 583, 409, 413, 412, 414, 411, 584, 585, 587, 588, 589, 594, 451, 595, 452, 586, 590, 591, 592, 593, 596, 597, 415, 453, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 630, 618, 620, 610, 611, 613, 614, 616, 612, 619, 617, 622, 623, 625, 626, 627, 628, 629, 621, 609, 624, 615, 631, 632, 608, 633, 634, 635, 636, 637, 638, 639, 640, 641, 643, 642, 645, 644, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 659, 658, 454, 660, 661, 422, 662, 663, 664, 665, 666, 455, 456, 457, 667, 668, 458, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 459, 460, 461, 679, 680, 681, 423, 682, 683, 684, 424, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 696, 697, 698, 695, 425, 462, 699, 463, 700, 701, 703, 702, 704, 707, 705, 706, 426, 464, 708, 465, 709, 710, 711, 712, 416, 419, 466, 713, 714, 715, 717, 718, 719, 716, 720, 467, 721, 722, 723, 468, 724, 725, 469, 726, 728, 418, 470, 727, 472, 729, 730, 471, 473, 731, 732, 733, 734, 735, 737, 736, 738, 739, 740, 474, 741, 742, 743, 744, 745, 475, 746, 747, 748, 749, 750, 751, 752, 427, 753, 754, 755, 758, 756, 757, 759, 760, 761, 762, 763, 769, 767, 768, 766, 764, 765, 476, 771, 770, 772, 774, 775, 776, 777, 478, 479, 778, 779, 482, 781, 780, 782, 785, 783, 784, 477, 787, 786, 788, 791, 789, 790, 480, 792, 481, 793, 794, 795, 398: 408},
		{389, 389},
		// 10
		{388, 388, 797},
		{4: 428, 483, 484, 485, 486, 487, 488, 430, 489, 490, 429, 491, 492, 493, 494, 495, 431, 432, 496, 497, 498, 499, 500, 501, 433, 434, 502, 503, 504, 505, 506, 435, 436, 507, 509, 508, 437, 510, 511, 512, 514, 513, 515, 438, 439, 516, 517, 518, 519, 520, 521, 523, 524, 525, 526, 528, 527, 440, 529, 530, 531, 532, 533, 534, 535, 537, 536, 538, 441, 539, 540, 541, 542, 543, 544, 545, 546, 547, 442, 548, 549, 550, 569, 522, 443, 551, 553, 554, 552, 555, 556, 557, 420, 558, 559, 560, 417, 561, 562, 563, 564, 565, 567, 566, 571, 421, 572, 573, 574, 444, 445, 446, 568, 570, 773, 575, 576, 577, 578, 580, 579, 447, 448, 581, 449, 450, 582, 583, 796, 413, 412, 414, 411, 584, 585, 587, 588, 589, 594, 451, 595, 452, 586, 590, 591, 592, 593, 596, 597, 415, 453, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 630, 618, 620, 610, 611, 613, 614, 616, 612, 619, 617, 622, 623, 625, 626, 627, 628, 629, 621, 609, 624, 615, 631, 632, 608, 633, 634, 635, 636, 637, 638, 639, 640, 641, 643, 642, 645, 644, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 659, 658, 454, 660, 661, 422, 662, 663, 664, 665, 666, 455, 456, 457, 667, 668, 458, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 459, 460, 461, 679, 680, 681, 423, 682, 683, 684, 424, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 696, 697, 698, 695, 425, 462, 699, 463, 700, 701, 703, 702, 704, 707, 705, 706, 426, 464, 708, 465, 709, 710, 711, 712, 416, 419, 466, 713, 714, 715, 717, 718, 719, 716, 720, 467, 721, 722, 723, 468, 724, 725, 469, 726, 728, 418, 470, 727, 472, 729, 730, 471, 473, 731, 732, 733, 734, 735, 737, 736, 738, 739, 740, 474, 741, 742, 743, 744, 745, 475, 746, 747, 748, 749, 750, 751, 752, 427, 753, 754, 755, 758, 756, 757, 759, 760, 761, 762, 763, 769, 767, 768, 766, 764, 765, 476, 771, 770, 772, 774, 775, 776, 777, 478, 479, 778, 779, 482, 781, 780, 782, 785, 783, 784, 477, 787, 786, 788, 791, 789, 790, 480, 792, 481, 793, 794, 795},
		{385, 385, 385},
		{384, 384, 384},
		{383, 383, 383},
		// 15
		{382, 382, 382},
		{381, 381, 381},
		{380, 380, 380},
		{379, 379, 379},
		{378, 378, 378},
		// 20
		{377, 377, 377},
		{376, 376, 376},
		{375, 375, 375},
		{374, 374, 374},
		{373, 373, 373},
		// 25
		{372, 372, 372},
		{371, 371, 371},
		{370, 370, 370},
		{369, 369, 369},
		{368, 368, 368},
		// 30
		{367, 367, 367},
		{366, 366, 366},
		{365, 365, 365},
		{364, 364, 364},
		{363, 363, 363},
		// 35
		{362, 362, 362},
		{361, 361, 361},
		{360, 360, 360},
		{359, 359, 359},
		{358, 358, 358},
		// 40
		{357, 357, 357},
		{356, 356, 356},
		{355, 355, 355},
		{354, 354, 354},
		{353, 353, 353},
		// 45
		{352, 352, 352},
		{351, 351, 351},
		{350, 350, 350},
		{349, 349, 349},
		{348, 348, 348},
		// 50
		{347, 347, 347},
		{346, 346, 346},
		{345, 345, 345},
		{344, 344, 344},
		{343, 343, 343},
		// 55
		{342, 342, 342},
		{341, 341, 341},
		{340, 340, 340},
		{339, 339, 339},
		{338, 338, 338},
		// 60
		{337, 337, 337},
		{336, 336, 336},
		{335, 335, 335},
		{334, 334, 334},
		{333, 333, 333},
		// 65
		{332, 332, 332},
		{331, 331, 331},
		{330, 330, 330},
		{329, 329, 329},
		{328, 328, 328},
		// 70
		{327, 327, 327},
		{326, 326, 326},
		{325, 325, 325},
		{324, 324, 324},
		{323, 323, 323},
		// 75
		{322, 322, 322},
		{321, 321, 321},
		{320, 320, 320},
		{319, 319, 319},
		{318, 318, 318},
		// 80
		{317, 317, 317},
		{316, 316, 316},
		{315, 315, 315},
		{314, 314, 314},
		{313, 313, 313},
		// 85
		{312, 312, 312},
		{311, 311, 311},
		{310, 310, 310},
		{309, 309, 309},
		{308, 308, 308},
		// 90
		{307, 307, 307},
		{306, 306, 306},
		{305, 305, 305},
		{304, 304, 304},
		{303, 303, 303},
		// 95
		{302, 302, 302},
		{301, 301, 301},
		{300, 300, 300},
		{299, 299, 299},
		{298, 298, 298},
		// 100
		{297, 297, 297},
		{296, 296, 296},
		{295, 295, 295},
		{294, 294, 294},
		{293, 293, 293},
		// 105
		{292, 292, 292},
		{291, 291, 291},
		{290, 290, 290},
		{289, 289, 289},
		{288, 288, 288},
		// 110
		{287, 287, 287},
		{286, 286, 286},
		{285, 285, 285},
		{284, 284, 284},
		{283, 283, 283},
		// 115
		{282, 282, 282},
		{281, 281, 281},
		{280, 280, 280},
		{279, 279, 279},
		{278, 278, 278},
		// 120
		{277, 277, 277},
		{276, 276, 276},
		{275, 275, 275},
		{274, 274, 274},
		{273, 273, 273},
		// 125
		{272, 272, 272},
		{271, 271, 271},
		{270, 270, 270},
		{269, 269, 269},
		{268, 268, 268},
		// 130
		{267, 267, 267},
		{266, 266, 266},
		{265, 265, 265},
		{264, 264, 264},
		{263, 263, 263},
		// 135
		{262, 262, 262},
		{261, 261, 261},
		{260, 260, 260},
		{259, 259, 259},
		{258, 258, 258},
		// 140
		{257, 257, 257},
		{256, 256, 256},
		{255, 255, 255},
		{254, 254, 254},
		{253, 253, 253},
		// 145
		{252, 252, 252},
		{251, 251, 251},
		{250, 250, 250},
		{249, 249, 249},
		{248, 248, 248},
		// 150
		{247, 247, 247},
		{246, 246, 246},
		{245, 245, 245},
		{244, 244, 244},
		{243, 243, 243},
		// 155
		{242, 242, 242},
		{241, 241, 241},
		{240, 240, 240},
		{239, 239, 239},
		{238, 238, 238},
		// 160
		{237, 237, 237},
		{236, 236, 236},
		{235, 235, 235},
		{234, 234, 234},
		{233, 233, 233},
		// 165
		{232, 232, 232},
		{231, 231, 231},
		{230, 230, 230},
		{229, 229, 229},
		{228, 228, 228},
		// 170
		{227, 227, 227},
		{226, 226, 226},
		{225, 225, 225},
		{224, 224, 224},
		{223, 223, 223},
		// 175
		{222, 222, 222},
		{221, 221, 221},
		{220, 220, 220},
		{219, 219, 219},
		{218, 218, 218},
		// 180
		{217, 217, 217},
		{216, 216, 216},
		{215, 215, 215},
		{214, 214, 214},
		{213, 213, 213},
		// 185
		{212, 212, 212},
		{211, 211, 211},
		{210, 210, 210},
		{209, 209, 209},
		{208, 208, 208},
		// 190
		{207, 207, 207},
		{206, 206, 206},
		{205, 205, 205},
		{204, 204, 204},
		{203, 203, 203},
		// 195
		{202, 202, 202},
		{201, 201, 201},
		{200, 200, 200},
		{199, 199, 199},
		{198, 198, 198},
		// 200
		{197, 197, 197},
		{196, 196, 196},
		{195, 195, 195},
		{194, 194, 194},
		{193, 193, 193},
		// 205
		{192, 192, 192},
		{191, 191, 191},
		{190, 190, 190},
		{189, 189, 189},
		{188, 188, 188},
		// 210
		{187, 187, 187},
		{186, 186, 186},
		{185, 185, 185},
		{184, 184, 184},
		{183, 183, 183},
		// 215
		{182, 182, 182},
		{181, 181, 181},
		{180, 180, 180},
		{179, 179, 179},
		{178, 178, 178},
		// 220
		{177, 177, 177},
		{176, 176, 176},
		{175, 175, 175},
		{174, 174, 174},
		{173, 173, 173},
		// 225
		{172, 172, 172},
		{171, 171, 171},
		{170, 170, 170},
		{169, 169, 169},
		{168, 168, 168},
		// 230
		{167, 167, 167},
		{166, 166, 166},
		{165, 165, 165},
		{164, 164, 164},
		{163, 163, 163},
		// 235
		{162, 162, 162},
		{161, 161, 161},
		{160, 160, 160},
		{159, 159, 159},
		{158, 158, 158},
		// 240
		{157, 157, 157},
		{156, 156, 156},
		{155, 155, 155},
		{154, 154, 154},
		{153, 153, 153},
		// 245
		{152, 152, 152},
		{151, 151, 151},
		{150, 150, 150},
		{149, 149, 149},
		{148, 148, 148},
		// 250
		{147, 147, 147},
		{146, 146, 146},
		{145, 145, 145},
		{144, 144, 144},
		{143, 143, 143},
		// 255
		{142, 142, 142},
		{141, 141, 141},
		{140, 140, 140},
		{139, 139, 139},
		{138, 138, 138},
		// 260
		{137, 137, 137},
		{136, 136, 136},
		{135, 135, 135},
		{134, 134, 134},
		{133, 133, 133},
		// 265
		{132, 132, 132},
		{131, 131, 131},
		{130, 130, 130},
		{129, 129, 129},
		{128, 128, 128},
		// 270
		{127, 127, 127},
		{126, 126, 126},
		{125, 125, 125},
		{124, 124, 124},
		{123, 123, 123},
		// 275
		{122, 122, 122},
		{121, 121, 121},
		{120, 120, 120},
		{119, 119, 119},
		{118, 118, 118},
		// 280
		{117, 117, 117},
		{116, 116, 116},
		{115, 115, 115},
		{114, 114, 114},
		{113, 113, 113},
		// 285
		{112, 112, 112},
		{111, 111, 111},
		{110, 110, 110},
		{109, 109, 109},
		{108, 108, 108},
		// 290
		{107, 107, 107},
		{106, 106, 106},
		{105, 105, 105},
		{104, 104, 104},
		{103, 103, 103},
		// 295
		{102, 102, 102},
		{101, 101, 101},
		{100, 100, 100},
		{99, 99, 99},
		{98, 98, 98},
		// 300
		{97, 97, 97},
		{96, 96, 96},
		{95, 95, 95},
		{94, 94, 94},
		{93, 93, 93},
		// 305
		{92, 92, 92},
		{91, 91, 91},
		{90, 90, 90},
		{89, 89, 89},
		{88, 88, 88},
		// 310
		{87, 87, 87},
		{86, 86, 86},
		{85, 85, 85},
		{84, 84, 84},
		{83, 83, 83},
		// 315
		{82, 82, 82},
		{81, 81, 81},
		{80, 80, 80},
		{79, 79, 79},
		{78, 78, 78},
		// 320
		{77, 77, 77},
		{76, 76, 76},
		{75, 75, 75},
		{74, 74, 74},
		{73, 73, 73},
		// 325
		{72, 72, 72},
		{71, 71, 71},
		{70, 70, 70},
		{69, 69, 69},
		{68, 68, 68},
		// 330
		{67, 67, 67},
		{66, 66, 66},
		{65, 65, 65},
		{64, 64, 64},
		{63, 63, 63},
		// 335
		{62, 62, 62},
		{61, 61, 61},
		{60, 60, 60},
		{59, 59, 59},
		{58, 58, 58},
		// 340
		{57, 57, 57},
		{56, 56, 56},
		{55, 55, 55},
		{54, 54, 54},
		{53, 53, 53},
		// 345
		{52, 52, 52},
		{51, 51, 51},
		{50, 50, 50},
		{49, 49, 49},
		{48, 48, 48},
		// 350
		{47, 47, 47},
		{46, 46, 46},
		{45, 45, 45},
		{44, 44, 44},
		{43, 43, 43},
		// 355
		{42, 42, 42},
		{41, 41, 41},
		{40, 40, 40},
		{39, 39, 39},
		{38, 38, 38},
		// 360
		{37, 37, 37},
		{36, 36, 36},
		{35, 35, 35},
		{34, 34, 34},
		{33, 33, 33},
		// 365
		{32, 32, 32},
		{31, 31, 31},
		{30, 30, 30},
		{29, 29, 29},
		{28, 28, 28},
		// 370
		{27, 27, 27},
		{26, 26, 26},
		{25, 25, 25},
		{24, 24, 24},
		{23, 23, 23},
		// 375
		{22, 22, 22},
		{21, 21, 21},
		{20, 20, 20},
		{19, 19, 19},
		{18, 18, 18},
		// 380
		{17, 17, 17},
		{16, 16, 16},
		{15, 15, 15},
		{14, 14, 14},
		{13, 13, 13},
		// 385
		{12, 12, 12},
		{11, 11, 11},
		{10, 10, 10},
		{9, 9, 9},
		{8, 8, 8},
		// 390
		{7, 7, 7},
		{6, 6, 6},
		{5, 5, 5},
		{4, 4, 4},
		{3, 3, 3},
		// 395
		{2, 2, 2},
		{1, 1, 1},
		{386, 386},
		{4: 428, 483, 484, 485, 486, 487, 488, 430, 489, 490, 429, 491, 492, 493, 494, 495, 431, 432, 496, 497, 498, 499, 500, 501, 433, 434, 502, 503, 504, 505, 506, 435, 436, 507, 509, 508, 437, 510, 511, 512, 514, 513, 515, 438, 439, 516, 517, 518, 519, 520, 521, 523, 524, 525, 526, 528, 527, 440, 529, 530, 531, 532, 533, 534, 535, 537, 536, 538, 441, 539, 540, 541, 542, 543, 544, 545, 546, 547, 442, 548, 549, 550, 569, 522, 443, 551, 553, 554, 552, 555, 556, 557, 420, 558, 559, 560, 417, 561, 562, 563, 564, 565, 567, 566, 571, 421, 572, 573, 574, 444, 445, 446, 568, 570, 773, 575, 576, 577, 578, 580, 579, 447, 448, 581, 449, 450, 582, 583, 798, 413, 412, 414, 411, 584, 585, 587, 588, 589, 594, 451, 595, 452, 586, 590, 591, 592, 593, 596, 597, 415, 453, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 630, 618, 620, 610, 611, 613, 614, 616, 612, 619, 617, 622, 623, 625, 626, 627, 628, 629, 621, 609, 624, 615, 631, 632, 608, 633, 634, 635, 636, 637, 638, 639, 640, 641, 643, 642, 645, 644, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 659, 658, 454, 660, 661, 422, 662, 663, 664, 665, 666, 455, 456, 457, 667, 668, 458, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 459, 460, 461, 679, 680, 681, 423, 682, 683, 684, 424, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 696, 697, 698, 695, 425, 462, 699, 463, 700, 701, 703, 702, 704, 707, 705, 706, 426, 464, 708, 465, 709, 710, 711, 712, 416, 419, 466, 713, 714, 715, 717, 718, 719, 716, 720, 467, 721, 722, 723, 468, 724, 725, 469, 726, 728, 418, 470, 727, 472, 729, 730, 471, 473, 731, 732, 733, 734, 735, 737, 736, 738, 739, 740, 474, 741, 742, 743, 744, 745, 475, 746, 747, 748, 749, 750, 751, 752, 427, 753, 754, 755, 758, 756, 757, 759, 760, 761, 762, 763, 769, 767, 768, 766, 764, 765, 476, 771, 770, 772, 774, 775, 776, 777, 478, 479, 778, 779, 482, 781, 780, 782, 785, 783, 784, 477, 787, 786, 788, 791, 789, 790, 480, 792, 481, 793, 794, 795},
		{387, 387},
		// 400
		{1: 801},
		{3: 394},
		{803, 3: 393, 393: 802},
		{3: 395},
		{3: 392},
	}
)

var MySQLDebug = 0

type MySQLLexer interface {
	Lex(lval *MySQLSymType) int
	Error(s string)
}

type MySQLLexerEx interface {
	MySQLLexer
	Reduced(rule, state int, lval *MySQLSymType) bool
}

func MySQLSymName(c int) (s string) {
	x, ok := MySQLXLAT[c]
	if ok {
		return MySQLSymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func MySQLlex1(yylex MySQLLexer, lval *MySQLSymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = MySQLEofCode
	}
	if MySQLDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", MySQLSymName(n), n, n, lval)
	}
	return n
}

func MySQLParse(yylex MySQLLexer) int {
	const yyError = 484

	yyEx, _ := yylex.(MySQLLexerEx)
	var yyn int
	var yylval MySQLSymType
	var yyVAL MySQLSymType
	yyS := make([]MySQLSymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if MySQLDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]MySQLSymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = MySQLlex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = MySQLXLAT[yychar]; !ok {
			yyxchar = len(MySQLSymNames) // > tab width
		}
	}
	if MySQLDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := MySQLParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += MySQLTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if MySQLDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if MySQLDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if MySQLDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", MySQLSymName(yychar), yystate)
			}
			msg, ok := MySQLXErrors[MySQLXError{yystate, yyxchar}]
			if !ok {
				msg, ok = MySQLXErrors[MySQLXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = MySQLXErrors[MySQLXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = MySQLXErrors[MySQLXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := MySQLTokenLiteralStrings[yychar]
				if ls == "" {
					ls = MySQLSymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := MySQLParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + MySQLTabOfs
					if yyn > 0 { // hit
						if MySQLDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if MySQLDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if MySQLDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if MySQLDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", MySQLSymName(yychar))
			}
			if yychar == MySQLEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := MySQLReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]MySQLSymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(MySQLParseTab[yyS[yyp].yys][x]) + MySQLTabOfs
	/* reduction by production r */
	if MySQLDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, MySQLSymNames[x], yystate)
	}

	switch r {
	case 2:
		{
			setParseTree(yylex, nil)
		}
	case 3:
		{
			setParseTree(yylex, yyS[yypt-0].statement)
		}
	case 5:
		{
			setParseTree(yylex, yyS[yypt-1].statement)
		}
	case 8:
		{
			yyVAL.statement = yyS[yypt-0].statement
		}
	case 9:
		{
			yyVAL.statement = yyS[yypt-0].statement
		}
	case 10:
		{
			yyVAL.statement = &AlterTableStmt{Table: yyS[yypt-0].table}
		}
	case 11:
		{
			yyVAL.table = &TableIdent{Name: yyS[yypt-0].bytes}
		}
	case 12:
		{
			yyVAL.table = &TableIdent{Schema: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 13:
		{
			yyVAL.table = &TableIdent{Name: yyS[yypt-0].bytes}
		}
	case 14:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 15:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 16:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 17:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 18:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 19:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 20:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 21:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 22:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 23:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 24:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 25:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 26:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 27:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 28:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 29:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 30:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 31:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 32:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 33:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 34:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 35:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 36:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 37:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 38:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 39:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 40:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 41:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 42:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 43:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 44:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 45:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 46:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 47:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 48:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 49:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 50:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 51:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 52:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 53:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 54:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 55:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 56:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 57:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 58:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 59:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 60:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 61:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 62:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 65:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 67:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 68:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 69:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 70:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 71:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 72:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 73:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 74:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 76:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 77:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 78:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 79:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 80:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 81:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 82:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 83:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 84:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 85:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 86:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 87:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 88:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 89:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 90:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 91:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 92:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 93:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 94:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 95:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 96:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 97:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 98:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 99:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 100:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 101:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 102:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 103:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 104:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 105:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 106:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 107:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 108:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 109:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 110:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 111:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 112:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 113:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 114:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 115:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 116:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 117:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 118:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 119:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 120:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 121:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 122:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 123:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 124:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 125:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 126:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 127:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 128:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 129:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 130:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 131:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 132:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 133:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 134:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 135:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 136:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 137:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 138:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 139:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 140:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 141:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 142:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 143:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 144:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 145:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 146:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 147:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 148:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 149:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 150:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 151:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 152:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 153:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 154:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 155:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 156:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 157:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 158:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 159:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 160:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 161:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 162:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 163:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 164:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 165:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 166:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 167:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 168:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 169:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 170:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 171:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 172:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 173:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 174:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 175:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 176:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 177:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 178:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 179:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 180:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 181:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 182:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 183:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 184:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 185:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 186:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 187:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 188:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 189:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 190:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 191:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 192:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 193:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 194:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 195:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 196:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 197:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 198:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 199:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 200:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 201:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 202:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 203:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 204:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 205:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 206:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 207:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 208:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 209:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 210:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 211:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 212:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 213:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 214:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 215:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 216:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 217:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 218:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 219:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 220:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 221:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 222:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 223:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 224:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 225:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 226:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 227:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 228:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 229:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 230:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 231:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 232:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 233:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 234:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 235:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 236:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 237:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 238:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 239:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 240:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 241:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 242:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 243:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 244:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 245:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 246:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 247:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 248:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 249:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 250:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 251:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 252:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 253:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 254:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 255:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 256:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 257:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 258:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 259:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 260:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 261:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 262:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 263:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 264:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 265:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 266:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 267:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 268:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 269:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 270:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 271:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 272:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 273:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 274:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 275:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 276:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 277:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 278:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 279:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 280:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 281:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 282:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 283:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 284:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 285:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 286:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 287:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 288:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 289:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 290:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 291:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 292:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 293:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 294:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 295:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 296:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 297:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 298:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 299:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 300:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 301:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 302:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 303:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 304:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 305:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 306:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 307:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 308:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 309:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 310:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 311:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 312:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 313:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 314:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 315:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 316:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 317:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 318:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 319:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 320:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 321:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 322:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 323:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 324:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 325:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 326:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 327:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 328:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 329:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 330:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 331:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 332:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 333:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 334:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 335:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 336:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 337:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 338:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 339:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 340:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 341:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 342:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 343:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 344:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 345:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 346:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 347:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 348:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 349:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 350:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 351:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 352:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 353:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 354:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 355:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 356:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 357:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 358:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 359:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 360:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 361:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 362:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 363:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 364:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 365:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 366:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 367:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 368:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 369:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 370:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 371:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 372:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 373:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 374:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 375:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 376:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 377:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 378:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 379:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 380:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 381:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 382:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 383:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 384:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 385:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 386:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 387:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 388:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 389:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 390:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 391:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 392:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 393:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 394:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 395:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 396:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 397:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 398:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 399:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 400:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 401:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 402:
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
