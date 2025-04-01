CREATE TABLE IF NOT EXISTS "Users"
(
    user_id serial NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    email_id character varying NOT NULL,
    user_password character varying NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    PRIMARY KEY (user_id)
);