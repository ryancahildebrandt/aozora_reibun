-- db.json is derived from the aozora bunko corpus DuckDB file.
-- from the aozora bunko corpus run: COPY (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\.|。|？|\?|!|！')) AS text FROM texts) TO 'db.json' (ARRAY);

-- these commands should be run on a sqlite db in the data directory
DROP TABLE texts;
CREATE TABLE texts(text TEXT);
INSERT INTO texts SELECT JSON_EXTRACT(VALUE, '$.text') FROM JSON_EACH(READFILE('db.json'));

DROP TABLE vtexts;
CREATE VIRTUAL TABLE vtexts USING FTS5(text);
INSERT INTO vtexts(text) SELECT text FROM texts;
