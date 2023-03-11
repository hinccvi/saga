BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "customer" (
  id  UUID  DEFAULT uuid_generate_v4(),
  name  VARCHAR(25) NOT NULL,
  credit_limit  NUMERIC(10,2) DEFAULT '0' NOT NULL,
  created_at    timestamp WITHOUT TIME ZONE NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
  PRIMARY KEY (id)
);
ALTER TABLE "customer" REPLICA IDENTITY USING INDEX customer_pkey;

CREATE TABLE IF NOT EXISTS "order" (
  id  UUID  DEFAULT uuid_generate_v4(),
  customer_id UUID NOT NULL,
  order_total NUMERIC(10,2) DEFAULT '0' NOT NULL,
  state VARCHAR(25) NOT NULL,
  PRIMARY KEY (id)
);

COMMIT;
