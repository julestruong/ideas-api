-- Table: public."user"

-- DROP TABLE public."user";

CREATE TABLE public."user"
(
  id SERIAL PRIMARY KEY,
  email character varying,
  firstname character varying,
  lastname character varying,
  CONSTRAINT user_email_key UNIQUE (email)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public."user"
  OWNER TO postgres;

