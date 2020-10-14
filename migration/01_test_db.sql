--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0
-- Dumped by pg_dump version 13.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: test; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE test WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';


ALTER DATABASE test OWNER TO postgres;

\connect test

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: new_user(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.new_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
        NEW.uuid = CONCAT('9', TO_CHAR(CURRENT_TIMESTAMP,'YYYWWDDDHH24MISSUS'), RIGHT(CAST(NEW.id AS VARCHAR), 3));
        NEW.created_at = now();
        END
$$;


ALTER FUNCTION public.new_user() OWNER TO postgres;

--
-- Name: update_user(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        NEW.modified_at = now();
        RETURN NEW;
        END;
    $$;


ALTER FUNCTION public.update_user() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    uuid character varying NOT NULL,
    nick_name character varying NOT NULL,
    login character varying NOT NULL,
    password character varying NOT NULL,
    rule bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone,
    blocked_at timestamp without time zone,
    blocked boolean
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: users_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_id_uindex ON public.users USING btree (id);


--
-- Name: users_login_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_login_uindex ON public.users USING btree (login);


--
-- Name: users_nick_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_nick_name_uindex ON public.users USING btree (nick_name);


--
-- Name: users_uuid_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_uuid_uindex ON public.users USING btree (uuid);


--
-- Name: users newuser; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER newuser BEFORE INSERT ON public.users FOR EACH ROW EXECUTE FUNCTION public.new_user();


--
-- Name: users updateuser; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updateuser BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_user();


--
-- PostgreSQL database dump complete
--

