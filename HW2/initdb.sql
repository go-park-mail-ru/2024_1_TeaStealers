CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "pg_stat_statements";
CREATE EXTENSION IF NOT EXISTS "auto_explain";

SET auto_explain.log_analyze = ON;
SET auto_explain.log_min_duration = '100ms';
SET auto_explain.log_timing = ON;

CREATE USER serv_acc WITH PASSWORD 'randomPasswordSec';
CREATE ROLE serv_acc_role WITH LOGIN;
GRANT USAGE ON SCHEMA public TO serv_acc_role;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO serv_acc_role;
GRANT serv_acc_role TO serv_acc;