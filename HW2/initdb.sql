CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "pg_stat_statements";
CREATE EXTENSION IF NOT EXISTS "auto_explain";

SET auto_explain.log_analyze = ON;
SET auto_explain.log_min_duration = '100ms';
SET auto_explain.log_timing = ON;

