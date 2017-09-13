-- Table: public."idea"

-- DROP TABLE public.idea;

CREATE TABLE public.idea
(
  id SERIAL PRIMARY KEY,
  email character varying NOT NULL,
  body character varying NOT NULL,
  week character varying NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(), 
  CONSTRAINT idea_user_email_week_key UNIQUE (email, week)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.idea
  OWNER TO postgres;

