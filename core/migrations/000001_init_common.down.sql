BEGIN;

DELETE DOMAIN resource_tags;
DELETE DOMAIN resource_description;
DELETE DOMAIN resource_name;
DELETE DOMAIN resource_key;
DELETE DOMAIN resource_id;

DELETE EXTENSION "pgcrypto";

END;
