-- Table: public."idea"

-- DROP TABLE public.idea;

CREATE TABLE public.idea
(
  id SERIAL PRIMARY KEY,
  email character varying,
  body character varying,
  created_at timestamp DEFAULT NOW(), 
  CONSTRAINT idea_user_email_key UNIQUE (email)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.idea
  OWNER TO postgres;

