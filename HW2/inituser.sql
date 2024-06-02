CREATE USER serv_acc WITH PASSWORD 'randomPasswordSec';
CREATE ROLE serv_acc_role WITH LOGIN;
GRANT USAGE ON SCHEMA public TO serv_acc_role;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO serv_acc_role;
GRANT serv_acc_role TO serv_acc;