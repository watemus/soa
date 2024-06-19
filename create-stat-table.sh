../ch/clickhouse client --host localhost --port 9000 --query "\
CREATE TABLE IF NOT EXISTS kafka_likes ( \
    username String, \
    task_id UInt64 \
) ENGINE = Kafka('kafka:9092', 'likes', 'likes', 'JSONEachRow');\
"
../ch/clickhouse client --host localhost --port 9000 --query "\
CREATE TABLE IF NOT EXISTS target_likes ( \
    username String, \
    task_id UInt64 \
) ENGINE = MergeTree()\
ORDER BY task_id\
"
../ch/clickhouse client --host localhost --port 9000 --query "\
CREATE MATERIALIZED VIEW IF NOT EXISTS kafka_likes_to_target \
TO target_likes AS \
SELECT * FROM kafka_likes;\
"



../ch/clickhouse client --host localhost --port 9000 --query "\
CREATE TABLE IF NOT EXISTS kafka_views ( \
    username String, \
    task_id UInt64 \
) ENGINE = Kafka('kafka:9092', 'views', 'views', 'JSONEachRow');\
"

../ch/clickhouse client --host localhost --port 9000 --query "\
CREATE TABLE IF NOT EXISTS target_views ( \
    username String, \
    task_id UInt64 \
) ENGINE = MergeTree()\
ORDER BY task_id\
"

../ch/clickhouse client --host localhost --port 9000 --query "\
CREATE MATERIALIZED VIEW IF NOT EXISTS kafka_views_to_target \
TO target_views AS \
SELECT * FROM kafka_views;\
