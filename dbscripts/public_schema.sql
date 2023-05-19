SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET client_min_messages = warning;
SET row_security = off;
CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;
SET search_path = public, pg_catalog;
SET default_tablespace = '';

CREATE TABLE users(
  id            uuid NOT NULL DEFAULT uuid_generate_v1mc(),
  first_name    TEXT NOT NULL,
  last_name     TEXT NOT NULL,
  username      TEXT NOT NULL,
  user_password TEXT NOT NULL,
  user_role     TEXT NOT NULL
);

CREATE TABLE apartments(
  id           uuid NOT NULL DEFAULT uuid_generate_v1mc(),
  size         NUMBER(10, 2) NOT NULL,
  room_numbers int NOT NULL,
  user_id uuid NOT NULL,
  CONSTRAINT   apartments_pk PRIMARY KEY (id),
  CONSTRAINT   fk_apartment_owner_or_tenant_id FOREIGN KEY (user_id)
    REFERENCES users(id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
);

CREATE TABLE services(
  id uuid     NOT NULL DEFAULT uuid_generate_v1mc(),
  name        TEXT NOT NULL,
  description TEXT NOT NULL,
  m_amount    TEXT NOT NULL,
  cost        NUMBER(10, 2) NOT NULL
);

CREATE TABLE bills(
  id uuid        NOT NULL DEFAULT uuid_generate_v1mc(),
  "from"         DATE NOT NULL,
  until          DATE NOT NULL,
  spend_amount   TEXT NOT NULL,
  payment_amount TEXT NOT NULL,
  paid           BOOLEAN NOT NULL DEFAULT false,
  apartment_id   uuid NOT NULL,
  service_id     uuid NOT NULL,
  created_at     timestamp NOT NULL DEFAULT NOW(),
  updated_at     timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT     fk_bill_owner_or_apartment_id FOREIGN KEY (apartment_id)
    REFERENCES   apartments(id),
  CONSTRAINT     fk_bill_service_id FOREIGN KEY (service_id)
    REFERENCES   services(id)
);

CREATE TABLE payments(
  id          uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
  bill_id     uuid UNIQUE NOT NULL,
  user_id     uuid NOT NULL,
  created_at  timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT  payment_bill_id FOREIGN KEY (bill_id) REFERENCES bills(id),
  CONSTRAINT  payment_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE apartment_services(
  apartment_id uuid NOT NULL,
  service_id   uuid NOT NULL,
  CONSTRAINT   subscribed_apartment FOREIGN KEY (apartment_id) 
    REFERENCES apartments(id),
  CONSTRAINT   subscribed_service FOREIGN KEY  (service_id) 
    REFERENCES services(id),
  PRIMARY KEY  (apartment_id, service_id)
);