%{
    package sql
    import ()
%}

%union {
   bytes []byte
   statement Statement
   table     Table
}


%token<bytes>  ABORT_SYM                     /* INTERNAL (used in lex) */
%token<bytes>  ACCESSIBLE_SYM
%token<bytes>  ACCOUNT_SYM
%token<bytes>  ACTION                        /* SQL-2003-N */
%token<bytes>  ADD                           /* SQL-2003-R */
%token<bytes>  ADDDATE_SYM                   /* MYSQL-FUNC */
%token<bytes>  AFTER_SYM                     /* SQL-2003-N */
%token<bytes>  AGAINST
%token<bytes>  AGGREGATE_SYM
%token<bytes>  ALGORITHM_SYM
%token<bytes>  ALL                           /* SQL-2003-R */
%token<bytes>  ALTER                         /* SQL-2003-R */
%token<bytes>  ALWAYS_SYM
%token<bytes>  ANALYSE_SYM
%token<bytes>  ANALYZE_SYM
%token<bytes>  AND_AND_SYM                   /* OPERATOR */
%token<bytes>  AND_SYM                       /* SQL-2003-R */
%token<bytes>  ANY_SYM                       /* SQL-2003-R */
%token<bytes>  AS                            /* SQL-2003-R */
%token<bytes>  ASC                           /* SQL-2003-N */
%token<bytes>  ASCII_SYM                     /* MYSQL-FUNC */
%token<bytes>  ASENSITIVE_SYM                /* FUTURE-USE */
%token<bytes>  AT_SYM                        /* SQL-2003-R */
%token<bytes>  AUTOEXTEND_SIZE_SYM
%token<bytes>  AUTO_INC
%token<bytes>  AVG_ROW_LENGTH
%token<bytes>  AVG_SYM                       /* SQL-2003-N */
%token<bytes>  BACKUP_SYM
%token<bytes>  BEFORE_SYM                    /* SQL-2003-N */
%token<bytes>  BEGIN_SYM                     /* SQL-2003-R */
%token<bytes>  BETWEEN_SYM                   /* SQL-2003-R */
%token<bytes>  BIGINT_SYM                    /* SQL-2003-R */
%token<bytes>  BINARY_SYM                    /* SQL-2003-R */
%token<bytes>  BINLOG_SYM
%token<bytes>  BIN_NUM
%token<bytes>  BIT_AND                       /* MYSQL-FUNC */
%token<bytes>  BIT_OR                        /* MYSQL-FUNC */
%token<bytes>  BIT_SYM                       /* MYSQL-FUNC */
%token<bytes>  BIT_XOR                       /* MYSQL-FUNC */
%token<bytes>  BLOB_SYM                      /* SQL-2003-R */
%token<bytes>  BLOCK_SYM
%token<bytes>  BOOLEAN_SYM                   /* SQL-2003-R */
%token<bytes>  BOOL_SYM
%token<bytes>  BOTH                          /* SQL-2003-R */
%token<bytes>  BTREE_SYM
%token<bytes>  BY                            /* SQL-2003-R */
%token<bytes>  BYTE_SYM
%token<bytes>  CACHE_SYM
%token<bytes>  CALL_SYM                      /* SQL-2003-R */
%token<bytes>  CASCADE                       /* SQL-2003-N */
%token<bytes>  CASCADED                      /* SQL-2003-R */ %token<bytes>  CASE_SYM                      /* SQL-2003-R */
%token<bytes>  CAST_SYM                      /* SQL-2003-R */
%token<bytes>  CATALOG_NAME_SYM              /* SQL-2003-N */
%token<bytes>  CHAIN_SYM                     /* SQL-2003-N */
%token<bytes>  CHANGE
%token<bytes>  CHANGED
%token<bytes>  CHANNEL_SYM
%token<bytes>  CHARSET
%token<bytes>  CHAR_SYM                      /* SQL-2003-R */
%token<bytes>  CHECKSUM_SYM
%token<bytes>  CHECK_SYM                     /* SQL-2003-R */
%token<bytes>  CIPHER_SYM
%token<bytes>  CLASS_ORIGIN_SYM              /* SQL-2003-N */
%token<bytes>  CLIENT_SYM
%token<bytes>  CLOSE_SYM                     /* SQL-2003-R */
%token<bytes>  COALESCE                      /* SQL-2003-N */
%token<bytes>  CODE_SYM
%token<bytes>  COLLATE_SYM                   /* SQL-2003-R */
%token<bytes>  COLLATION_SYM                 /* SQL-2003-N */
%token<bytes>  COLUMNS
%token<bytes>  COLUMN_SYM                    /* SQL-2003-R */
%token<bytes>  COLUMN_FORMAT_SYM
%token<bytes>  COLUMN_NAME_SYM               /* SQL-2003-N */
%token<bytes>  COMMENT_SYM
%token<bytes>  COMMITTED_SYM                 /* SQL-2003-N */
%token<bytes>  COMMIT_SYM                    /* SQL-2003-R */
%token<bytes>  COMPACT_SYM
%token<bytes>  COMPLETION_SYM
%token<bytes>  COMPRESSED_SYM
%token<bytes>  COMPRESSION_SYM
%token<bytes>  ENCRYPTION_SYM
%token<bytes>  CONCURRENT
%token<bytes>  CONDITION_SYM                 /* SQL-2003-R, SQL-2008-R */
%token<bytes>  CONNECTION_SYM
%token<bytes>  CONSISTENT_SYM
%token<bytes>  CONSTRAINT                    /* SQL-2003-R */
%token<bytes>  CONSTRAINT_CATALOG_SYM        /* SQL-2003-N */
%token<bytes>  CONSTRAINT_NAME_SYM           /* SQL-2003-N */
%token<bytes>  CONSTRAINT_SCHEMA_SYM         /* SQL-2003-N */
%token<bytes>  CONTAINS_SYM                  /* SQL-2003-N */
%token<bytes>  CONTEXT_SYM
%token<bytes>  CONTINUE_SYM                  /* SQL-2003-R */
%token<bytes>  CONVERT_SYM                   /* SQL-2003-N */
%token<bytes>  COUNT_SYM                     /* SQL-2003-N */
%token<bytes>  CPU_SYM
%token<bytes>  CREATE                        /* SQL-2003-R */
%token<bytes>  CROSS                         /* SQL-2003-R */
%token<bytes>  CUBE_SYM                      /* SQL-2003-R */
%token<bytes>  CURDATE                       /* MYSQL-FUNC */
%token<bytes>  CURRENT_SYM                   /* SQL-2003-R */
%token<bytes>  CURRENT_USER                  /* SQL-2003-R */
%token<bytes>  CURSOR_SYM                    /* SQL-2003-R */
%token<bytes>  CURSOR_NAME_SYM               /* SQL-2003-N */
%token<bytes>  CURTIME                       /* MYSQL-FUNC */
%token<bytes>  DATABASE
%token<bytes>  DATABASES
%token<bytes>  DATAFILE_SYM
%token<bytes>  DATA_SYM                      /* SQL-2003-N */
%token<bytes>  DATETIME_SYM                  /* MYSQL */
%token<bytes>  DATE_ADD_INTERVAL             /* MYSQL-FUNC */
%token<bytes>  DATE_SUB_INTERVAL             /* MYSQL-FUNC */
%token<bytes>  DATE_SYM                      /* SQL-2003-R */
%token<bytes>  DAY_HOUR_SYM
%token<bytes>  DAY_MICROSECOND_SYM
%token<bytes>  DAY_MINUTE_SYM
%token<bytes>  DAY_SECOND_SYM
%token<bytes>  DAY_SYM                       /* SQL-2003-R */
%token<bytes>  DEALLOCATE_SYM                /* SQL-2003-R */
%token<bytes>  DECIMAL_NUM
%token<bytes>  DECIMAL_SYM                   /* SQL-2003-R */
%token<bytes>  DECLARE_SYM                   /* SQL-2003-R */
%token<bytes>  DEFAULT_SYM                   /* SQL-2003-R */
%token<bytes>  DEFAULT_AUTH_SYM              /* INTERNAL */
%token<bytes>  DEFINER_SYM
%token<bytes>  DELAYED_SYM
%token<bytes>  DELAY_KEY_WRITE_SYM
%token<bytes>  DELETE_SYM                    /* SQL-2003-R */
%token<bytes>  DESC                          /* SQL-2003-N */
%token<bytes>  DESCRIBE                      /* SQL-2003-R */
%token<bytes>  DES_KEY_FILE
%token<bytes>  DETERMINISTIC_SYM             /* SQL-2003-R */
%token<bytes>  DIAGNOSTICS_SYM               /* SQL-2003-N */
%token<bytes>  DIRECTORY_SYM
%token<bytes>  DISABLE_SYM
%token<bytes>  DISCARD
%token<bytes>  DISK_SYM
%token<bytes>  DISTINCT                      /* SQL-2003-R */
%token<bytes>  DIV_SYM
%token<bytes>  DOUBLE_SYM                    /* SQL-2003-R */
%token<bytes>  DO_SYM
%token<bytes>  DROP                          /* SQL-2003-R */
%token<bytes>  DUAL_SYM
%token<bytes>  DUMPFILE
%token<bytes>  DUPLICATE_SYM
%token<bytes>  DYNAMIC_SYM                   /* SQL-2003-R */
%token<bytes>  EACH_SYM                      /* SQL-2003-R */
%token<bytes>  ELSE                          /* SQL-2003-R */
%token<bytes>  ELSEIF_SYM
%token<bytes>  ENABLE_SYM
%token<bytes>  ENCLOSED
%token<bytes>  END                           /* SQL-2003-R */
%token<bytes>  ENDS_SYM
%token<bytes>  END_OF_INPUT                  /* INTERNAL */
%token<bytes>  ENGINES_SYM
%token<bytes>  ENGINE_SYM
%token<bytes>  ENUM_SYM                      /* MYSQL */
%token<bytes>  EQ                            /* OPERATOR */
%token<bytes>  EQUAL_SYM                     /* OPERATOR */
%token<bytes>  ERROR_SYM
%token<bytes>  ERRORS
%token<bytes>  ESCAPED
%token<bytes>  ESCAPE_SYM                    /* SQL-2003-R */
%token<bytes>  EVENTS_SYM
%token<bytes>  EVENT_SYM
%token<bytes>  EVERY_SYM                     /* SQL-2003-N */
%token<bytes>  EXCHANGE_SYM
%token<bytes>  EXECUTE_SYM                   /* SQL-2003-R */
%token<bytes>  EXISTS                        /* SQL-2003-R */
%token<bytes>  EXIT_SYM
%token<bytes>  EXPANSION_SYM
%token<bytes>  EXPIRE_SYM
%token<bytes>  EXPORT_SYM
%token<bytes>  EXTENDED_SYM
%token<bytes>  EXTENT_SIZE_SYM
%token<bytes>  EXTRACT_SYM                   /* SQL-2003-N */
%token<bytes>  FALSE_SYM                     /* SQL-2003-R */
%token<bytes>  FAST_SYM
%token<bytes>  FAULTS_SYM
%token<bytes>  FETCH_SYM                     /* SQL-2003-R */
%token<bytes>  FILE_SYM
%token<bytes>  FILE_BLOCK_SIZE_SYM
%token<bytes>  FILTER_SYM
%token<bytes>  FIRST_SYM                     /* SQL-2003-N */
%token<bytes>  FIXED_SYM
%token<bytes>  FLOAT_NUM
%token<bytes>  FLOAT_SYM                     /* SQL-2003-R */
%token<bytes>  FLUSH_SYM
%token<bytes>  FOLLOWS_SYM                  /* MYSQL */
%token<bytes>  FORCE_SYM
%token<bytes>  FOREIGN                       /* SQL-2003-R */
%token<bytes>  FOR_SYM                       /* SQL-2003-R */
%token<bytes>  FORMAT_SYM
%token<bytes>  FOUND_SYM                     /* SQL-2003-R */
%token<bytes>  FROM
%token<bytes>  FULL                          /* SQL-2003-R */
%token<bytes>  FULLTEXT_SYM
%token<bytes>  FUNCTION_SYM                  /* SQL-2003-R */
%token<bytes>  GE
%token<bytes>  GENERAL
%token<bytes>  GENERATED
%token<bytes>  GROUP_REPLICATION
%token<bytes>  GEOMETRYCOLLECTION_SYM        /* MYSQL */
%token<bytes>  GEOMETRY_SYM
%token<bytes>  GET_FORMAT                    /* MYSQL-FUNC */
%token<bytes>  GET_SYM                       /* SQL-2003-R */
%token<bytes>  GLOBAL_SYM                    /* SQL-2003-R */
%token<bytes>  GRANT                         /* SQL-2003-R */
%token<bytes>  GRANTS
%token<bytes>  GROUP_SYM                     /* SQL-2003-R */
%token<bytes>  GROUP_CONCAT_SYM
%token<bytes>  GT_SYM                        /* OPERATOR */
%token<bytes>  HANDLER_SYM
%token<bytes>  HASH_SYM
%token<bytes>  HAVING                        /* SQL-2003-R */
%token<bytes>  HELP_SYM
%token<bytes>  HEX_NUM
%token<bytes>  HIGH_PRIORITY
%token<bytes>  HOST_SYM
%token<bytes>  HOSTS_SYM
%token<bytes>  HOUR_MICROSECOND_SYM
%token<bytes>  HOUR_MINUTE_SYM
%token<bytes>  HOUR_SECOND_SYM
%token<bytes>  HOUR_SYM                      /* SQL-2003-R */
%token<bytes>  IDENT
%token<bytes>  IDENTIFIED_SYM
%token<bytes>  IDENT_QUOTED
%token<bytes>  IF
%token<bytes>  IGNORE_SYM
%token<bytes>  IGNORE_SERVER_IDS_SYM
%token<bytes>  IMPORT
%token<bytes>  INDEXES
%token<bytes>  INDEX_SYM
%token<bytes>  INFILE
%token<bytes>  INITIAL_SIZE_SYM
%token<bytes>  INNER_SYM                     /* SQL-2003-R */
%token<bytes>  INOUT_SYM                     /* SQL-2003-R */
%token<bytes>  INSENSITIVE_SYM               /* SQL-2003-R */
%token<bytes>  INSERT_SYM                    /* SQL-2003-R */
%token<bytes>  INSERT_METHOD
%token<bytes>  INSTANCE_SYM
%token<bytes>  INSTALL_SYM
%token<bytes>  INTERVAL_SYM                  /* SQL-2003-R */
%token<bytes>  INTO                          /* SQL-2003-R */
%token<bytes>  INT_SYM                       /* SQL-2003-R */
%token<bytes>  INVOKER_SYM
%token<bytes>  IN_SYM                        /* SQL-2003-R */
%token<bytes>  IO_AFTER_GTIDS                /* MYSQL, FUTURE-USE */
%token<bytes>  IO_BEFORE_GTIDS               /* MYSQL, FUTURE-USE */
%token<bytes>  IO_SYM
%token<bytes>  IPC_SYM
%token<bytes>  IS                            /* SQL-2003-R */
%token<bytes>  ISOLATION                     /* SQL-2003-R */
%token<bytes>  ISSUER_SYM
%token<bytes>  ITERATE_SYM
%token<bytes>  JOIN_SYM                      /* SQL-2003-R */
%token<bytes>  JSON_SEPARATOR_SYM            /* MYSQL */
%token<bytes>  JSON_SYM                      /* MYSQL */
%token<bytes>  KEYS
%token<bytes>  KEY_BLOCK_SIZE
%token<bytes>  KEY_SYM                       /* SQL-2003-N */
%token<bytes>  KILL_SYM
%token<bytes>  LANGUAGE_SYM                  /* SQL-2003-R */
%token<bytes>  LAST_SYM                      /* SQL-2003-N */
%token<bytes>  LE                            /* OPERATOR */
%token<bytes>  LEADING                       /* SQL-2003-R */
%token<bytes>  LEAVES
%token<bytes>  LEAVE_SYM
%token<bytes>  LEFT                          /* SQL-2003-R */
%token<bytes>  LESS_SYM
%token<bytes>  LEVEL_SYM
%token<bytes>  LEX_HOSTNAME
%token<bytes>  LIKE                          /* SQL-2003-R */
%token<bytes>  LIMIT
%token<bytes>  LINEAR_SYM
%token<bytes>  LINES
%token<bytes>  LINESTRING_SYM                /* MYSQL */
%token<bytes>  LIST_SYM
%token<bytes>  LOAD
%token<bytes>  LOCAL_SYM                     /* SQL-2003-R */
%token<bytes>  LOCATOR_SYM                   /* SQL-2003-N */
%token<bytes>  LOCKS_SYM
%token<bytes>  LOCK_SYM
%token<bytes>  LOGFILE_SYM
%token<bytes>  LOGS_SYM
%token<bytes>  LONGBLOB_SYM                  /* MYSQL */
%token<bytes>  LONGTEXT_SYM                  /* MYSQL */
%token<bytes>  LONG_NUM
%token<bytes>  LONG_SYM
%token<bytes>  LOOP_SYM
%token<bytes>  LOW_PRIORITY
%token<bytes>  LT                            /* OPERATOR */
%token<bytes>  MASTER_AUTO_POSITION_SYM
%token<bytes>  MASTER_BIND_SYM
%token<bytes>  MASTER_CONNECT_RETRY_SYM
%token<bytes>  MASTER_DELAY_SYM
%token<bytes>  MASTER_HOST_SYM
%token<bytes>  MASTER_LOG_FILE_SYM
%token<bytes>  MASTER_LOG_POS_SYM
%token<bytes>  MASTER_PASSWORD_SYM
%token<bytes>  MASTER_PORT_SYM
%token<bytes>  MASTER_RETRY_COUNT_SYM
%token<bytes>  MASTER_SERVER_ID_SYM
%token<bytes>  MASTER_SSL_CAPATH_SYM
%token<bytes>  MASTER_TLS_VERSION_SYM
%token<bytes>  MASTER_SSL_CA_SYM
%token<bytes>  MASTER_SSL_CERT_SYM
%token<bytes>  MASTER_SSL_CIPHER_SYM
%token<bytes>  MASTER_SSL_CRL_SYM
%token<bytes>  MASTER_SSL_CRLPATH_SYM
%token<bytes>  MASTER_SSL_KEY_SYM
%token<bytes>  MASTER_SSL_SYM
%token<bytes>  MASTER_SSL_VERIFY_SERVER_CERT_SYM
%token<bytes>  MASTER_SYM
%token<bytes>  MASTER_USER_SYM
%token<bytes>  MASTER_HEARTBEAT_PERIOD_SYM
%token<bytes>  MATCH                         /* SQL-2003-R */
%token<bytes>  MAX_CONNECTIONS_PER_HOUR
%token<bytes>  MAX_QUERIES_PER_HOUR
%token<bytes>  MAX_ROWS
%token<bytes>  MAX_SIZE_SYM
%token<bytes>  MAX_SYM                       /* SQL-2003-N */
%token<bytes>  MAX_UPDATES_PER_HOUR
%token<bytes>  MAX_USER_CONNECTIONS_SYM
%token<bytes>  MAX_VALUE_SYM                 /* SQL-2003-N */
%token<bytes>  MEDIUMBLOB_SYM                /* MYSQL */
%token<bytes>  MEDIUMINT_SYM                 /* MYSQL */
%token<bytes>  MEDIUMTEXT_SYM                /* MYSQL */
%token<bytes>  MEDIUM_SYM
%token<bytes>  MEMORY_SYM
%token<bytes>  MERGE_SYM                     /* SQL-2003-R */
%token<bytes>  MESSAGE_TEXT_SYM              /* SQL-2003-N */
%token<bytes>  MICROSECOND_SYM               /* MYSQL-FUNC */
%token<bytes>  MIGRATE_SYM
%token<bytes>  MINUTE_MICROSECOND_SYM
%token<bytes>  MINUTE_SECOND_SYM
%token<bytes>  MINUTE_SYM                    /* SQL-2003-R */
%token<bytes>  MIN_ROWS
%token<bytes>  MIN_SYM                       /* SQL-2003-N */
%token<bytes>  MODE_SYM
%token<bytes>  MODIFIES_SYM                  /* SQL-2003-R */
%token<bytes>  MODIFY_SYM
%token<bytes>  MOD_SYM                       /* SQL-2003-N */
%token<bytes>  MONTH_SYM                     /* SQL-2003-R */
%token<bytes>  MULTILINESTRING_SYM           /* MYSQL */
%token<bytes>  MULTIPOINT_SYM                /* MYSQL */
%token<bytes>  MULTIPOLYGON_SYM              /* MYSQL */
%token<bytes>  MUTEX_SYM
%token<bytes>  MYSQL_ERRNO_SYM
%token<bytes>  NAMES_SYM                     /* SQL-2003-N */
%token<bytes>  NAME_SYM                      /* SQL-2003-N */
%token<bytes>  NATIONAL_SYM                  /* SQL-2003-R */
%token<bytes>  NATURAL                       /* SQL-2003-R */
%token<bytes>  NCHAR_STRING
%token<bytes>  NCHAR_SYM                     /* SQL-2003-R */
%token<bytes>  NDBCLUSTER_SYM
%token<bytes>  NE                            /* OPERATOR */
%token<bytes>  NEG
%token<bytes>  NEVER_SYM
%token<bytes>  NEW_SYM                       /* SQL-2003-R */
%token<bytes>  NEXT_SYM                      /* SQL-2003-N */
%token<bytes>  NODEGROUP_SYM
%token<bytes>  NONE_SYM                      /* SQL-2003-R */
%token<bytes>  NOT2_SYM
%token<bytes>  NOT_SYM                       /* SQL-2003-R */
%token<bytes>  NOW_SYM
%token<bytes>  NO_SYM                        /* SQL-2003-R */
%token<bytes>  NO_WAIT_SYM
%token<bytes>  NO_WRITE_TO_BINLOG
%token<bytes>  NULL_SYM                      /* SQL-2003-R */
%token<bytes>  NUM
%token<bytes>  NUMBER_SYM                    /* SQL-2003-N */
%token<bytes>  NUMERIC_SYM                   /* SQL-2003-R */
%token<bytes>  NVARCHAR_SYM
%token<bytes>  OFFSET_SYM
%token<bytes>  ON_SYM                        /* SQL-2003-R */
%token<bytes>  ONE_SYM
%token<bytes>  ONLY_SYM                      /* SQL-2003-R */
%token<bytes>  OPEN_SYM                      /* SQL-2003-R */
%token<bytes>  OPTIMIZE
%token<bytes>  OPTIMIZER_COSTS_SYM
%token<bytes>  OPTIONS_SYM
%token<bytes>  OPTION                        /* SQL-2003-N */
%token<bytes>  OPTIONALLY
%token<bytes>  OR2_SYM
%token<bytes>  ORDER_SYM                     /* SQL-2003-R */
%token<bytes>  OR_OR_SYM                     /* OPERATOR */
%token<bytes>  OR_SYM                        /* SQL-2003-R */
%token<bytes>  OUTER
%token<bytes>  OUTFILE
%token<bytes>  OUT_SYM                       /* SQL-2003-R */
%token<bytes>  OWNER_SYM
%token<bytes>  PACK_KEYS_SYM
%token<bytes>  PAGE_SYM
%token<bytes>  PARAM_MARKER
%token<bytes>  PARSER_SYM
%token<bytes>  OBSOLETE_TOKEN_654            /* was: PARSE_GCOL_EXPR_SYM */
%token<bytes>  PARTIAL                       /* SQL-2003-N */
%token<bytes>  PARTITION_SYM                 /* SQL-2003-R */
%token<bytes>  PARTITIONS_SYM
%token<bytes>  PARTITIONING_SYM
%token<bytes>  PASSWORD
%token<bytes>  PHASE_SYM
%token<bytes>  PLUGIN_DIR_SYM                /* INTERNAL */
%token<bytes>  PLUGIN_SYM
%token<bytes>  PLUGINS_SYM
%token<bytes>  POINT_SYM
%token<bytes>  POLYGON_SYM                   /* MYSQL */
%token<bytes>  PORT_SYM
%token<bytes>  POSITION_SYM                  /* SQL-2003-N */
%token<bytes>  PRECEDES_SYM                  /* MYSQL */
%token<bytes>  PRECISION                     /* SQL-2003-R */
%token<bytes>  PREPARE_SYM                   /* SQL-2003-R */
%token<bytes>  PRESERVE_SYM
%token<bytes>  PREV_SYM
%token<bytes>  PRIMARY_SYM                   /* SQL-2003-R */
%token<bytes>  PRIVILEGES                    /* SQL-2003-N */
%token<bytes>  PROCEDURE_SYM                 /* SQL-2003-R */
%token<bytes>  PROCESS
%token<bytes>  PROCESSLIST_SYM
%token<bytes>  PROFILE_SYM
%token<bytes>  PROFILES_SYM
%token<bytes>  PROXY_SYM
%token<bytes>  PURGE
%token<bytes>  QUARTER_SYM
%token<bytes>  QUERY_SYM
%token<bytes>  QUICK
%token<bytes>  RANGE_SYM                     /* SQL-2003-R */
%token<bytes>  READS_SYM                     /* SQL-2003-R */
%token<bytes>  READ_ONLY_SYM
%token<bytes>  READ_SYM                      /* SQL-2003-N */
%token<bytes>  READ_WRITE_SYM
%token<bytes>  REAL_SYM                      /* SQL-2003-R */
%token<bytes>  REBUILD_SYM
%token<bytes>  RECOVER_SYM
%token<bytes>  REDOFILE_SYM
%token<bytes>  REDO_BUFFER_SIZE_SYM
%token<bytes>  REDUNDANT_SYM
%token<bytes>  REFERENCES                    /* SQL-2003-R */
%token<bytes>  REGEXP
%token<bytes>  RELAY
%token<bytes>  RELAYLOG_SYM
%token<bytes>  RELAY_LOG_FILE_SYM
%token<bytes>  RELAY_LOG_POS_SYM
%token<bytes>  RELAY_THREAD
%token<bytes>  RELEASE_SYM                   /* SQL-2003-R */
%token<bytes>  RELOAD
%token<bytes>  REMOVE_SYM
%token<bytes>  RENAME
%token<bytes>  REORGANIZE_SYM
%token<bytes>  REPAIR
%token<bytes>  REPEATABLE_SYM                /* SQL-2003-N */
%token<bytes>  REPEAT_SYM                    /* MYSQL-FUNC */
%token<bytes>  REPLACE_SYM                   /* MYSQL-FUNC */
%token<bytes>  REPLICATION
%token<bytes>  REPLICATE_DO_DB
%token<bytes>  REPLICATE_IGNORE_DB
%token<bytes>  REPLICATE_DO_TABLE
%token<bytes>  REPLICATE_IGNORE_TABLE
%token<bytes>  REPLICATE_WILD_DO_TABLE
%token<bytes>  REPLICATE_WILD_IGNORE_TABLE
%token<bytes>  REPLICATE_REWRITE_DB
%token<bytes>  REQUIRE_SYM
%token<bytes>  RESET_SYM
%token<bytes>  RESIGNAL_SYM                  /* SQL-2003-R */
%token<bytes>  RESOURCES
%token<bytes>  RESTORE_SYM
%token<bytes>  RESTRICT
%token<bytes>  RESUME_SYM
%token<bytes>  RETURNED_SQLSTATE_SYM         /* SQL-2003-N */
%token<bytes>  RETURNS_SYM                   /* SQL-2003-R */
%token<bytes>  RETURN_SYM                    /* SQL-2003-R */
%token<bytes>  REVERSE_SYM
%token<bytes>  REVOKE                        /* SQL-2003-R */
%token<bytes>  RIGHT                         /* SQL-2003-R */
%token<bytes>  ROLLBACK_SYM                  /* SQL-2003-R */
%token<bytes>  ROLLUP_SYM                    /* SQL-2003-R */
%token<bytes>  ROTATE_SYM
%token<bytes>  ROUTINE_SYM                   /* SQL-2003-N */
%token<bytes>  ROWS_SYM                      /* SQL-2003-R */
%token<bytes>  ROW_FORMAT_SYM
%token<bytes>  ROW_SYM                       /* SQL-2003-R */
%token<bytes>  ROW_COUNT_SYM                 /* SQL-2003-N */
%token<bytes>  RTREE_SYM
%token<bytes>  SAVEPOINT_SYM                 /* SQL-2003-R */
%token<bytes>  SCHEDULE_SYM
%token<bytes>  SCHEMA_NAME_SYM               /* SQL-2003-N */
%token<bytes>  SECOND_MICROSECOND_SYM
%token<bytes>  SECOND_SYM                    /* SQL-2003-R */
%token<bytes>  SECURITY_SYM                  /* SQL-2003-N */
%token<bytes>  SELECT_SYM                    /* SQL-2003-R */
%token<bytes>  SENSITIVE_SYM                 /* FUTURE-USE */
%token<bytes>  SEPARATOR_SYM
%token<bytes>  SERIALIZABLE_SYM              /* SQL-2003-N */
%token<bytes>  SERIAL_SYM
%token<bytes>  SESSION_SYM                   /* SQL-2003-N */
%token<bytes>  SERVER_SYM
%token<bytes>  SERVER_OPTIONS
%token<bytes>  SET_SYM                       /* SQL-2003-R */
%token<bytes>  SET_VAR
%token<bytes>  SHARE_SYM
%token<bytes>  SHIFT_LEFT                    /* OPERATOR */
%token<bytes>  SHIFT_RIGHT                   /* OPERATOR */
%token<bytes>  SHOW
%token<bytes>  SHUTDOWN
%token<bytes>  SIGNAL_SYM                    /* SQL-2003-R */
%token<bytes>  SIGNED_SYM
%token<bytes>  SIMPLE_SYM                    /* SQL-2003-N */
%token<bytes>  SLAVE
%token<bytes>  SLOW
%token<bytes>  SMALLINT_SYM                  /* SQL-2003-R */
%token<bytes>  SNAPSHOT_SYM
%token<bytes>  SOCKET_SYM
%token<bytes>  SONAME_SYM
%token<bytes>  SOUNDS_SYM
%token<bytes>  SOURCE_SYM
%token<bytes>  SPATIAL_SYM
%token<bytes>  SPECIFIC_SYM                  /* SQL-2003-R */
%token<bytes>  SQLEXCEPTION_SYM              /* SQL-2003-R */
%token<bytes>  SQLSTATE_SYM                  /* SQL-2003-R */
%token<bytes>  SQLWARNING_SYM                /* SQL-2003-R */
%token<bytes>  SQL_AFTER_GTIDS               /* MYSQL */
%token<bytes>  SQL_AFTER_MTS_GAPS            /* MYSQL */
%token<bytes>  SQL_BEFORE_GTIDS              /* MYSQL */
%token<bytes>  SQL_BIG_RESULT
%token<bytes>  SQL_BUFFER_RESULT
%token<bytes>  SQL_CACHE_SYM
%token<bytes>  SQL_CALC_FOUND_ROWS
%token<bytes>  SQL_NO_CACHE_SYM
%token<bytes>  SQL_SMALL_RESULT
%token<bytes>  SQL_SYM                       /* SQL-2003-R */
%token<bytes>  SQL_THREAD
%token<bytes>  SSL_SYM
%token<bytes>  STACKED_SYM                   /* SQL-2003-N */
%token<bytes>  STARTING
%token<bytes>  STARTS_SYM
%token<bytes>  START_SYM                     /* SQL-2003-R */
%token<bytes>  STATS_AUTO_RECALC_SYM
%token<bytes>  STATS_PERSISTENT_SYM
%token<bytes>  STATS_SAMPLE_PAGES_SYM
%token<bytes>  STATUS_SYM
%token<bytes>  STDDEV_SAMP_SYM               /* SQL-2003-N */
%token<bytes>  STD_SYM
%token<bytes>  STOP_SYM
%token<bytes>  STORAGE_SYM
%token<bytes>  STORED_SYM
%token<bytes>  STRAIGHT_JOIN
%token<bytes>  STRING_SYM
%token<bytes>  SUBCLASS_ORIGIN_SYM           /* SQL-2003-N */
%token<bytes>  SUBDATE_SYM
%token<bytes>  SUBJECT_SYM
%token<bytes>  SUBPARTITIONS_SYM
%token<bytes>  SUBPARTITION_SYM
%token<bytes>  SUBSTRING                     /* SQL-2003-N */
%token<bytes>  SUM_SYM                       /* SQL-2003-N */
%token<bytes>  SUPER_SYM
%token<bytes>  SUSPEND_SYM
%token<bytes>  SWAPS_SYM
%token<bytes>  SWITCHES_SYM
%token<bytes>  SYSDATE
%token<bytes>  TABLES
%token<bytes>  TABLESPACE_SYM
%token<bytes>  OBSOLETE_TOKEN_820            /* was: TABLE_REF_PRIORITY */
%token<bytes>  TABLE_SYM                     /* SQL-2003-R */
%token<bytes>  TABLE_CHECKSUM_SYM
%token<bytes>  TABLE_NAME_SYM                /* SQL-2003-N */
%token<bytes>  TEMPORARY                     /* SQL-2003-N */
%token<bytes>  TEMPTABLE_SYM
%token<bytes>  TERMINATED
%token<bytes>  TEXT_STRING
%token<bytes>  TEXT_SYM
%token<bytes>  THAN_SYM
%token<bytes>  THEN_SYM                      /* SQL-2003-R */
%token<bytes>  TIMESTAMP_SYM                 /* SQL-2003-R */
%token<bytes>  TIMESTAMP_ADD
%token<bytes>  TIMESTAMP_DIFF
%token<bytes>  TIME_SYM                      /* SQL-2003-R */
%token<bytes>  TINYBLOB_SYM                  /* MYSQL */
%token<bytes>  TINYINT_SYM                   /* MYSQL */
%token<bytes>  TINYTEXT_SYN                  /* MYSQL */
%token<bytes>  TO_SYM                        /* SQL-2003-R */
%token<bytes>  TRAILING                      /* SQL-2003-R */
%token<bytes>  TRANSACTION_SYM
%token<bytes>  TRIGGERS_SYM
%token<bytes>  TRIGGER_SYM                   /* SQL-2003-R */
%token<bytes>  TRIM                          /* SQL-2003-N */
%token<bytes>  TRUE_SYM                      /* SQL-2003-R */
%token<bytes>  TRUNCATE_SYM
%token<bytes>  TYPES_SYM
%token<bytes>  TYPE_SYM                      /* SQL-2003-N */
%token<bytes>  UDF_RETURNS_SYM
%token<bytes>  ULONGLONG_NUM
%token<bytes>  UNCOMMITTED_SYM               /* SQL-2003-N */
%token<bytes>  UNDEFINED_SYM
%token<bytes>  UNDERSCORE_CHARSET
%token<bytes>  UNDOFILE_SYM
%token<bytes>  UNDO_BUFFER_SIZE_SYM
%token<bytes>  UNDO_SYM                      /* FUTURE-USE */
%token<bytes>  UNICODE_SYM
%token<bytes>  UNINSTALL_SYM
%token<bytes>  UNION_SYM                     /* SQL-2003-R */
%token<bytes>  UNIQUE_SYM
%token<bytes>  UNKNOWN_SYM                   /* SQL-2003-R */
%token<bytes>  UNLOCK_SYM
%token<bytes>  UNSIGNED_SYM                  /* MYSQL */
%token<bytes>  UNTIL_SYM
%token<bytes>  UPDATE_SYM                    /* SQL-2003-R */
%token<bytes>  UPGRADE_SYM
%token<bytes>  USAGE                         /* SQL-2003-N */
%token<bytes>  USER                          /* SQL-2003-R */
%token<bytes>  USE_FRM
%token<bytes>  USE_SYM
%token<bytes>  USING                         /* SQL-2003-R */
%token<bytes>  UTC_DATE_SYM
%token<bytes>  UTC_TIMESTAMP_SYM
%token<bytes>  UTC_TIME_SYM
%token<bytes>  VALIDATION_SYM                /* MYSQL */
%token<bytes>  VALUES                        /* SQL-2003-R */
%token<bytes>  VALUE_SYM                     /* SQL-2003-R */
%token<bytes>  VARBINARY_SYM                 /* SQL-2008-R */
%token<bytes>  VARCHAR_SYM                   /* SQL-2003-R */
%token<bytes>  VARIABLES
%token<bytes>  VARIANCE_SYM
%token<bytes>  VARYING                       /* SQL-2003-R */
%token<bytes>  VAR_SAMP_SYM
%token<bytes>  VIEW_SYM                      /* SQL-2003-N */
%token<bytes>  VIRTUAL_SYM
%token<bytes>  WAIT_SYM
%token<bytes>  WARNINGS
%token<bytes>  WEEK_SYM
%token<bytes>  WEIGHT_STRING_SYM
%token<bytes>  WHEN_SYM                      /* SQL-2003-R */
%token<bytes>  WHERE                         /* SQL-2003-R */
%token<bytes>  WHILE_SYM
%token<bytes>  WITH                          /* SQL-2003-R */
%token<bytes>  WITH_CUBE_SYM                 /* INTERNAL */
%token<bytes>  WITH_ROLLUP_SYM               /* INTERNAL */
%token<bytes>  WITHOUT_SYM                   /* SQL-2003-R */
%token<bytes>  WORK_SYM                      /* SQL-2003-N */
%token<bytes>  WRAPPER_SYM
%token<bytes>  WRITE_SYM                     /* SQL-2003-N */
%token<bytes>  X509_SYM
%token<bytes>  XA_SYM
%token<bytes>  XID_SYM                       /* MYSQL */
%token<bytes>  XML_SYM
%token<bytes>  XOR
%token<bytes>  YEAR_MONTH_SYM
%token<bytes>  YEAR_SYM                      /* SQL-2003-R */
%token<bytes>  ZEROFILL_SYM                  /* MYSQL */

/*
   Tokens from MySQL 8.0
*/
%token<bytes>  JSON_UNQUOTED_SEPARATOR_SYM   /* MYSQL */
%token<bytes>  PERSIST_SYM
%token<bytes>  ROLE_SYM                      /* SQL-1999-R */
%token<bytes>  ADMIN_SYM                     /* SQL-1999-R */
%token<bytes>  INVISIBLE_SYM
%token<bytes>  VISIBLE_SYM
%token<bytes>  EXCEPT_SYM                    /* SQL-1999-R */
%token<bytes>  COMPONENT_SYM                 /* MYSQL */
%token<bytes>  GRAMMAR_SELECTOR_EXPR         /* synthetic token: starts single expr. */
%token<bytes>  GRAMMAR_SELECTOR_GCOL       /* synthetic token: starts generated col. */
%token<bytes>  GRAMMAR_SELECTOR_PART      /* synthetic token: starts partition expr. */

%right UNIQUE_SYM KEY_SYM

%left CONDITIONLESS_JOIN
%left   JOIN_SYM INNER_SYM CROSS STRAIGHT_JOIN NATURAL LEFT RIGHT ON_SYM USING
%left   SET_VAR
%left   OR_OR_SYM OR_SYM OR2_SYM
%left   XOR
%left   AND_SYM AND_AND_SYM
%left   BETWEEN_SYM CASE_SYM WHEN_SYM THEN_SYM ELSE
%left   EQ EQUAL_SYM GE GT_SYM LE LT NE IS LIKE REGEXP IN_SYM
%left   '|'
%left   '&'
%left   SHIFT_LEFT SHIFT_RIGHT
%left   '-' '+'
%left   '*' '/' '%' DIV_SYM MOD_SYM
%left   '^'
%left   NEG '~'
%right  NOT_SYM NOT2_SYM
%right  BINARY_SYM COLLATE_SYM
%left  INTERVAL_SYM
%left SUBQUERY_AS_EXPR
%left '(' ')'

%left EMPTY_FROM_CLAUSE
%right INTO

%type <statement> simple_statement_or_begin simple_statement 
%type <statement> alter
%type <table>  table_ident
%type <bytes> ident IDENT_sys ident_keyword role_or_label_keyword role_or_ident_keyword label_keyword ident_or_text TEXT_STRING_sys


%%

start_entry:
           sql_statement
           ;

sql_statement:
             END_OF_INPUT
             {
                setParseTree(yylex, nil)
             }
             | simple_statement_or_begin
             {
                setParseTree(yylex, $1)
             }
             ';'
             opt_end_of_input
             |simple_statement_or_begin END_OF_INPUT
             {
                setParseTree(yylex, $1)
             }
             ;

opt_end_of_input:
                /* empty */
                | END_OF_INPUT
                ;

simple_statement_or_begin:
                         simple_statement { $$ = $1}
                         ;

simple_statement:
                alter {$$ = $1}
                ;


/*
** Alter table
*/

alter:
     ALTER TABLE_SYM table_ident
     {
         $$ = &AlterTableStmt{Table: $3}
     }
     ;



















table_ident:
           ident
           {
             $$ = &TableIdent{Name:$1}; 
           }
           | ident '.' ident
           {
             $$ = &TableIdent{Schema:$1,  Name:$3}; 
           }
           |'.' ident
           {
             $$  = &TableIdent{Name:$2}; 
           }
           ;


ident:
     IDENT_sys { $$ = $1 }
     |  ident_keyword { $$ = $1 }
     ;


IDENT_sys:
         IDENT { $$ = $1 }
         | IDENT_QUOTED { $$ = $1 }
         ;

ident_keyword:
             label_keyword         { $$ = $1}
             | role_or_ident_keyword {$$ = $1}
             | EXECUTE_SYM           {$$ = $1}
             | SHUTDOWN              {$$ = $1}
             ;
label_keyword:
             role_or_label_keyword    { $$=$1 }
             | EVENT_SYM                { $$=$1 }
             | FILE_SYM                 { $$=$1 }
             | NONE_SYM                 { $$=$1 }
             | PROCESS                  { $$=$1 }
             | PROXY_SYM                { $$=$1 }
             | RELOAD                   { $$=$1}
             | REPLICATION              { $$=$1 }
             | SUPER_SYM                { $$=$1 }
             ;
role_or_ident_keyword:
          ACCOUNT_SYM           { $$=$1 }
        | ASCII_SYM             { $$=$1 }
        | ALWAYS_SYM            { $$=$1 }
        | BACKUP_SYM            { $$=$1 }
        | BEGIN_SYM             { $$=$1 }
        | BYTE_SYM              { $$=$1 }
        | CACHE_SYM             { $$=$1 }
        | CHARSET               { $$=$1 }
        | CHECKSUM_SYM          { $$=$1 }
        | CLOSE_SYM             { $$=$1 }
        | COMMENT_SYM           { $$=$1 }
        | COMMIT_SYM            { $$=$1 }
        | CONTAINS_SYM          { $$=$1 }
        | DEALLOCATE_SYM        { $$=$1 }
        | DO_SYM                { $$=$1 }
        | END                   { $$=$1 }
        | FLUSH_SYM             { $$=$1 }
        | FOLLOWS_SYM           { $$=$1 }
        | FORMAT_SYM            { $$=$1 }
        | GROUP_REPLICATION     { $$=$1 }
        | HANDLER_SYM           { $$=$1 }
        | HELP_SYM              { $$=$1 }
        | HOST_SYM              { $$=$1 }
        | INSTALL_SYM           { $$=$1 }
        | INVISIBLE_SYM         { $$=$1 }
        | LANGUAGE_SYM          { $$=$1 }
        | NO_SYM                { $$=$1 }
        | OPEN_SYM              { $$=$1 }
        | OPTIONS_SYM           { $$=$1 }
        | OWNER_SYM             { $$=$1 }
        | PARSER_SYM            { $$=$1 }
        | PORT_SYM              { $$=$1 }
        | PRECEDES_SYM          { $$=$1 }
        | PREPARE_SYM           { $$=$1 }
        | REMOVE_SYM            { $$=$1 }
        | REPAIR                { $$=$1 }
        | RESET_SYM             { $$=$1 }
        | RESTORE_SYM           { $$=$1 }
        | ROLLBACK_SYM          { $$=$1 }
        | SAVEPOINT_SYM         { $$=$1 }
        | SECURITY_SYM          { $$=$1 }
        | SERVER_SYM            { $$=$1 }
        | SIGNED_SYM            { $$=$1 }
        | SOCKET_SYM            { $$=$1 }
        | SLAVE                 { $$=$1 }
        | SONAME_SYM            { $$=$1 }
        | START_SYM             { $$=$1 }
        | STOP_SYM              { $$=$1 }
        | TRUNCATE_SYM          { $$=$1 }
        | VISIBLE_SYM           { $$=$1 }
        | UNICODE_SYM           { $$=$1 }
        | UNINSTALL_SYM         { $$=$1 }
        | WRAPPER_SYM           { $$=$1 }
        | XA_SYM                { $$=$1 }
        | UPGRADE_SYM           { $$=$1 }
        ;
role_or_label_keyword:
          ACTION                   { $$=$1 }
        | ADDDATE_SYM              { $$=$1 }
        | AFTER_SYM                { $$=$1 }
        | AGAINST                  { $$=$1 }
        | AGGREGATE_SYM            { $$=$1 }
        | ALGORITHM_SYM            { $$=$1 }
        | ANALYSE_SYM              { $$=$1 }
        | ANY_SYM                  { $$=$1 }
        | AT_SYM                   { $$=$1 }
        | AUTO_INC                 { $$=$1 }
        | AUTOEXTEND_SIZE_SYM      { $$=$1 }
        | AVG_ROW_LENGTH           { $$=$1 }
        | AVG_SYM                  { $$=$1 }
        | BINLOG_SYM               { $$=$1 }
        | BIT_SYM                  { $$=$1 }
        | BLOCK_SYM                { $$=$1 }
        | BOOL_SYM                 { $$=$1 }
        | BOOLEAN_SYM              { $$=$1 }
        | BTREE_SYM                { $$=$1 }
        | CASCADED                 { $$=$1 }
        | CATALOG_NAME_SYM         { $$=$1 }
        | CHAIN_SYM                { $$=$1 }
        | CHANGED                  { $$=$1 }
        | CHANNEL_SYM              { $$=$1 }
        | CIPHER_SYM               { $$=$1 }
        | CLIENT_SYM               { $$=$1 }
        | CLASS_ORIGIN_SYM         { $$=$1 }
        | COALESCE                 { $$=$1 }
        | CODE_SYM                 { $$=$1 }
        | COLLATION_SYM            { $$=$1 }
        | COLUMN_NAME_SYM          { $$=$1 }
        | COLUMN_FORMAT_SYM        { $$=$1 }
        | COLUMNS                  { $$=$1 }
        | COMMITTED_SYM            { $$=$1 }
        | COMPACT_SYM              { $$=$1 }
        | COMPLETION_SYM           { $$=$1 }
        | COMPONENT_SYM            { $$=$1 }
        | COMPRESSED_SYM           { $$=$1 }
        | COMPRESSION_SYM          { $$=$1 }
        | ENCRYPTION_SYM           { $$=$1 }
        | CONCURRENT               { $$=$1 }
        | CONNECTION_SYM           { $$=$1 }
        | CONSISTENT_SYM           { $$=$1 }
        | CONSTRAINT_CATALOG_SYM   { $$=$1 }
        | CONSTRAINT_SCHEMA_SYM    { $$=$1 }
        | CONSTRAINT_NAME_SYM      { $$=$1 }
        | CONTEXT_SYM              { $$=$1 }
        | CPU_SYM                  { $$=$1 }
        | CUBE_SYM                 { $$=$1 }
        /*
          Although a reserved keyword in SQL:2003 (and :2008),
          not reserved in MySQL per WL#2111 specification.
        */
        | CURRENT_SYM              { $$=$1 }
        | CURSOR_NAME_SYM          { $$=$1 }
        | DATA_SYM                 { $$=$1 }
        | DATAFILE_SYM             { $$=$1 }
        | DATETIME_SYM             { $$=$1 }
        | DATE_SYM                 { $$=$1 }
        | DAY_SYM                  { $$=$1 }
        | DEFAULT_AUTH_SYM         { $$=$1 }
        | DEFINER_SYM              { $$=$1 }
        | DELAY_KEY_WRITE_SYM      { $$=$1 }
        | DES_KEY_FILE             { $$=$1 }
        | DIAGNOSTICS_SYM          { $$=$1 }
        | DIRECTORY_SYM            { $$=$1 }
        | DISABLE_SYM              { $$=$1 }
        | DISCARD                  { $$=$1 }
        | DISK_SYM                 { $$=$1 }
        | DUMPFILE                 { $$=$1 }
        | DUPLICATE_SYM            { $$=$1 }
        | DYNAMIC_SYM              { $$=$1 }
        | ENDS_SYM                 { $$=$1 }
        | ENUM_SYM                 { $$=$1 }
        | ENGINE_SYM               { $$=$1 }
        | ENGINES_SYM              { $$=$1 }
        | ERROR_SYM                { $$=$1 }
        | ERRORS                   { $$=$1 }
        | ESCAPE_SYM               { $$=$1 }
        | EVENTS_SYM               { $$=$1 }
        | EVERY_SYM                { $$=$1 }
        | EXCHANGE_SYM             { $$=$1 }
        | EXPANSION_SYM            { $$=$1 }
        | EXPIRE_SYM               { $$=$1 }
        | EXPORT_SYM               { $$=$1 }
        | EXTENDED_SYM             { $$=$1 }
        | EXTENT_SIZE_SYM          { $$=$1 }
        | FAULTS_SYM               { $$=$1 }
        | FAST_SYM                 { $$=$1 }
        | FOUND_SYM                { $$=$1 }
        | ENABLE_SYM               { $$=$1 }
        | FULL                     { $$=$1 }
        | FILE_BLOCK_SIZE_SYM      { $$=$1 }
        | FILTER_SYM               { $$=$1 }
        | FIRST_SYM                { $$=$1 }
        | FIXED_SYM                { $$=$1 }
        | GENERAL                  { $$=$1 }
        | GEOMETRY_SYM             { $$=$1 }
        | GEOMETRYCOLLECTION_SYM   { $$=$1 }
        | GET_FORMAT               { $$=$1 }
        | GRANTS                   { $$=$1 }
        | GLOBAL_SYM               { $$=$1 }
        | HASH_SYM                 { $$=$1 }
        | HOSTS_SYM                { $$=$1 }
        | HOUR_SYM                 { $$=$1 }
        | IDENTIFIED_SYM           { $$=$1 }
        | IGNORE_SERVER_IDS_SYM    { $$=$1 }
        | INVOKER_SYM              { $$=$1 }
        | IMPORT                   { $$=$1 }
        | INDEXES                  { $$=$1 }
        | INITIAL_SIZE_SYM         { $$=$1 }
        | IO_SYM                   { $$=$1 }
        | IPC_SYM                  { $$=$1 }
        | ISOLATION                { $$=$1 }
        | ISSUER_SYM               { $$=$1 }
        | INSERT_METHOD            { $$=$1 }
        | INSTANCE_SYM             { $$=$1 }
        | JSON_SYM                 { $$=$1 }
        | KEY_BLOCK_SIZE           { $$=$1 }
        | LAST_SYM                 { $$=$1 }
        | LEAVES                   { $$=$1 }
        | LESS_SYM                 { $$=$1 }
        | LEVEL_SYM                { $$=$1 }
        | LINESTRING_SYM           { $$=$1 }
        | LIST_SYM                 { $$=$1 }
        | LOCAL_SYM                { $$=$1 }
        | LOCKS_SYM                { $$=$1 }
        | LOGFILE_SYM              { $$=$1 }
        | LOGS_SYM                 { $$=$1 }
        | MAX_ROWS                 { $$=$1 }
        | MASTER_SYM               { $$=$1 }
        | MASTER_HEARTBEAT_PERIOD_SYM { $$=$1 }
        | MASTER_HOST_SYM          { $$=$1 }
        | MASTER_PORT_SYM          { $$=$1 }
        | MASTER_LOG_FILE_SYM      { $$=$1 }
        | MASTER_LOG_POS_SYM       { $$=$1 }
        | MASTER_USER_SYM          { $$=$1 }
        | MASTER_PASSWORD_SYM      { $$=$1 }
        | MASTER_SERVER_ID_SYM     { $$=$1 }
        | MASTER_CONNECT_RETRY_SYM { $$=$1 }
        | MASTER_RETRY_COUNT_SYM   { $$=$1 }
        | MASTER_DELAY_SYM         { $$=$1 }
        | MASTER_SSL_SYM           { $$=$1 }
        | MASTER_SSL_CA_SYM        { $$=$1 }
        | MASTER_SSL_CAPATH_SYM    { $$=$1 }
        | MASTER_TLS_VERSION_SYM   { $$=$1 }
        | MASTER_SSL_CERT_SYM      { $$=$1 }
        | MASTER_SSL_CIPHER_SYM    { $$=$1 }
        | MASTER_SSL_CRL_SYM       { $$=$1 }
        | MASTER_SSL_CRLPATH_SYM   { $$=$1 }
        | MASTER_SSL_KEY_SYM       { $$=$1 }
        | MASTER_AUTO_POSITION_SYM { $$=$1 }
        | MAX_CONNECTIONS_PER_HOUR { $$=$1 }
        | MAX_QUERIES_PER_HOUR     { $$=$1 }
        | MAX_SIZE_SYM             { $$=$1 }
        | MAX_UPDATES_PER_HOUR     { $$=$1 }
        | MAX_USER_CONNECTIONS_SYM { $$=$1 }
        | MEDIUM_SYM               { $$=$1 }
        | MEMORY_SYM               { $$=$1 }
        | MERGE_SYM                { $$=$1 }
        | MESSAGE_TEXT_SYM         { $$=$1 }
        | MICROSECOND_SYM          { $$=$1 }
        | MIGRATE_SYM              { $$=$1 }
        | MINUTE_SYM               { $$=$1 }
        | MIN_ROWS                 { $$=$1 }
        | MODIFY_SYM               { $$=$1 }
        | MODE_SYM                 { $$=$1 }
        | MONTH_SYM                { $$=$1 }
        | MULTILINESTRING_SYM      { $$=$1 }
        | MULTIPOINT_SYM           { $$=$1 }
        | MULTIPOLYGON_SYM         { $$=$1 }
        | MUTEX_SYM                { $$=$1 }
        | MYSQL_ERRNO_SYM          { $$=$1 }
        | NAME_SYM                 { $$=$1 }
        | NAMES_SYM                { $$=$1 }
        | NATIONAL_SYM             { $$=$1 }
        | NCHAR_SYM                { $$=$1 }
        | NDBCLUSTER_SYM           { $$=$1 }
        | NEVER_SYM                { $$=$1 }
        | NEXT_SYM                 { $$=$1 }
        | NEW_SYM                  { $$=$1 }
        | NO_WAIT_SYM              { $$=$1 }
        | NODEGROUP_SYM            { $$=$1 }
        | NUMBER_SYM               { $$=$1 }
        | NVARCHAR_SYM             { $$=$1 }
        | OFFSET_SYM               { $$=$1 }
        | ONE_SYM                  { $$=$1 }
        | ONLY_SYM                 { $$=$1 }
        | PACK_KEYS_SYM            { $$=$1 }
        | PAGE_SYM                 { $$=$1 }
        | PARTIAL                  { $$=$1 }
        | PARTITIONING_SYM         { $$=$1 }
        | PARTITIONS_SYM           { $$=$1 }
        | PASSWORD                 { $$=$1 }
        | PHASE_SYM                { $$=$1 }
        | PLUGIN_DIR_SYM           { $$=$1 }
        | PLUGIN_SYM               { $$=$1 }
        | PLUGINS_SYM              { $$=$1 }
        | POINT_SYM                { $$=$1 }
        | POLYGON_SYM              { $$=$1 }
        | PRESERVE_SYM             { $$=$1 }
        | PREV_SYM                 { $$=$1 }
        | PRIVILEGES               { $$=$1 }
        | PROCESSLIST_SYM          { $$=$1 }
        | PROFILE_SYM              { $$=$1 }
        | PROFILES_SYM             { $$=$1 }
        | QUARTER_SYM              { $$=$1 }
        | QUERY_SYM                { $$=$1 }
        | QUICK                    { $$=$1 }
        | READ_ONLY_SYM            { $$=$1 }
        | REBUILD_SYM              { $$=$1 }
        | RECOVER_SYM              { $$=$1 }
        | REDO_BUFFER_SIZE_SYM     { $$=$1 }
        | REDOFILE_SYM             { $$=$1 }
        | REDUNDANT_SYM            { $$=$1 }
        | RELAY                    { $$=$1 }
        | RELAYLOG_SYM             { $$=$1 }
        | RELAY_LOG_FILE_SYM       { $$=$1 }
        | RELAY_LOG_POS_SYM        { $$=$1 }
        | RELAY_THREAD             { $$=$1 }
        | REORGANIZE_SYM           { $$=$1 }
        | REPEATABLE_SYM           { $$=$1 }
        | REPLICATE_DO_DB          { $$=$1 }
        | REPLICATE_IGNORE_DB      { $$=$1 }
        | REPLICATE_DO_TABLE       { $$=$1 }
        | REPLICATE_IGNORE_TABLE   { $$=$1 }
        | REPLICATE_WILD_DO_TABLE  { $$=$1 }
        | REPLICATE_WILD_IGNORE_TABLE { $$=$1}
        | REPLICATE_REWRITE_DB     { $$=$1 }
        | RESOURCES                { $$=$1 }
        | RESUME_SYM               { $$=$1 }
        | RETURNED_SQLSTATE_SYM    { $$=$1 }
        | RETURNS_SYM              { $$=$1 }
        | REVERSE_SYM              { $$=$1 }
        | ROLLUP_SYM               { $$=$1 }
        | ROTATE_SYM               { $$=$1 }
        | ROUTINE_SYM              { $$=$1 }
        | ROWS_SYM                 { $$=$1 }
        | ROW_COUNT_SYM            { $$=$1 }
        | ROW_FORMAT_SYM           { $$=$1 }
        | ROW_SYM                  { $$=$1 }
        | RTREE_SYM                { $$=$1 }
        | SCHEDULE_SYM             { $$=$1 }
        | SCHEMA_NAME_SYM          { $$=$1 }
        | SECOND_SYM               { $$=$1 }
        | SERIAL_SYM               { $$=$1 }
        | SERIALIZABLE_SYM         { $$=$1 }
        | SESSION_SYM              { $$=$1 }
        | SIMPLE_SYM               { $$=$1 }
        | SHARE_SYM                { $$=$1 }
        | SLOW                     { $$=$1 }
        | SNAPSHOT_SYM             { $$=$1 }
        | SOUNDS_SYM               { $$=$1 }
        | SOURCE_SYM               { $$=$1 }
        | SQL_AFTER_GTIDS          { $$=$1 }
        | SQL_AFTER_MTS_GAPS       { $$=$1 }
        | SQL_BEFORE_GTIDS         { $$=$1 }
        | SQL_CACHE_SYM            { $$=$1 }
        | SQL_BUFFER_RESULT        { $$=$1 }
        | SQL_NO_CACHE_SYM         { $$=$1 }
        | SQL_THREAD               { $$=$1 }
        | STACKED_SYM              { $$=$1 }
        | STARTS_SYM               { $$=$1 }
        | STATS_AUTO_RECALC_SYM    { $$=$1 }
        | STATS_PERSISTENT_SYM     { $$=$1 }
        | STATS_SAMPLE_PAGES_SYM   { $$=$1 }
        | STATUS_SYM               { $$=$1 }
        | STORAGE_SYM              { $$=$1 }
        | STRING_SYM               { $$=$1 }
        | SUBCLASS_ORIGIN_SYM      { $$=$1 }
        | SUBDATE_SYM              { $$=$1 }
        | SUBJECT_SYM              { $$=$1 }
        | SUBPARTITION_SYM         { $$=$1 }
        | SUBPARTITIONS_SYM        { $$=$1 }
        | SUSPEND_SYM              { $$=$1 }
        | SWAPS_SYM                { $$=$1 }
        | SWITCHES_SYM             { $$=$1 }
        | TABLE_NAME_SYM           { $$=$1 }
        | TABLES                   { $$=$1 }
        | TABLE_CHECKSUM_SYM       { $$=$1 }
        | TABLESPACE_SYM           { $$=$1 }
        | TEMPORARY                { $$=$1 }
        | TEMPTABLE_SYM            { $$=$1 }
        | TEXT_SYM                 { $$=$1 }
        | THAN_SYM                 { $$=$1 }
        | TRANSACTION_SYM          { $$=$1 }
        | TRIGGERS_SYM             { $$=$1 }
        | TIMESTAMP_SYM            { $$=$1 }
        | TIMESTAMP_ADD            { $$=$1 }
        | TIMESTAMP_DIFF           { $$=$1 }
        | TIME_SYM                 { $$=$1 }
        | TYPES_SYM                { $$=$1 }
        | TYPE_SYM                 { $$=$1 }
        | UDF_RETURNS_SYM          { $$=$1 }
        | FUNCTION_SYM             { $$=$1 }
        | UNCOMMITTED_SYM          { $$=$1 }
        | UNDEFINED_SYM            { $$=$1 }
        | UNDO_BUFFER_SIZE_SYM     { $$=$1 }
        | UNDOFILE_SYM             { $$=$1 }
        | UNKNOWN_SYM              { $$=$1 }
        | UNTIL_SYM                { $$=$1 }
        | USER                     { $$=$1 }
        | USE_FRM                  { $$=$1 }
        | VALIDATION_SYM           { $$=$1 }
        | VARIABLES                { $$=$1 }
        | VIEW_SYM                 { $$=$1 }
        | VALUE_SYM                { $$=$1 }
        | WARNINGS                 { $$=$1 }
        | WAIT_SYM                 { $$=$1 }
        | WEEK_SYM                 { $$=$1 }
        | WITHOUT_SYM              { $$=$1 }
        | WORK_SYM                 { $$=$1 }
        | WEIGHT_STRING_SYM        { $$=$1 }
        | X509_SYM                 { $$=$1 }
        | XID_SYM                  { $$=$1 }
        | XML_SYM                  { $$=$1 }
        | YEAR_SYM                 { $$=$1 }
        ;




ident_or_text:
             ident           { $$=$1}
             | TEXT_STRING_sys { $$=$1}
             | LEX_HOSTNAME { $$=$1;}
             ;

TEXT_STRING_sys:
               TEXT_STRING { $$ = $1} 
               ;
