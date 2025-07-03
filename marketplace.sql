--
-- PostgreSQL database dump
--

-- Dumped from database version 15.13 (Debian 15.13-0+deb12u1)
-- Dumped by pg_dump version 15.13 (Debian 15.13-0+deb12u1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cart; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.cart (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    product_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    quantity integer DEFAULT 1 NOT NULL
);


--
-- Name: cart_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.cart_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: cart_id_seq1; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.cart_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: cart_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.cart_id_seq1 OWNED BY public.cart.id;


--
-- Name: invoice; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.invoice (
    id bigint NOT NULL,
    price integer NOT NULL,
    details text NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: invoice_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.invoice_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: invoice_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.invoice_id_seq OWNED BY public.invoice.id;


--
-- Name: invoice_order_list; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.invoice_order_list (
    id bigint NOT NULL,
    invoice_id bigint NOT NULL,
    order_id bigint NOT NULL
);


--
-- Name: invoice_order_list_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.invoice_order_list_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: invoice_order_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.invoice_order_list_id_seq OWNED BY public.invoice_order_list.id;


--
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: order; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."order" (
    id bigint DEFAULT nextval('public.order_id_seq'::regclass) NOT NULL,
    buyer_id bigint NOT NULL,
    seller_id bigint NOT NULL,
    product_id bigint NOT NULL,
    price integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    quantity bigint NOT NULL
);


--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: product; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.product (
    id bigint DEFAULT nextval('public.product_id_seq'::regclass) NOT NULL,
    name character varying(200) NOT NULL,
    description text,
    photo character varying(200) NOT NULL,
    price integer NOT NULL,
    stock integer NOT NULL,
    slug character varying(200) NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."user" (
    id bigint DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    username character varying(100) NOT NULL,
    password character varying(200) NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    user_role_id bigint NOT NULL,
    is_active smallint DEFAULT '1'::smallint NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


--
-- Name: user_module; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_module (
    id bigint NOT NULL,
    name character varying(200) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: user_module_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_module_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


--
-- Name: user_module_id_seq1; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_module_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_module_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_module_id_seq1 OWNED BY public.user_module.id;


--
-- Name: user_role; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_role (
    id bigint NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: user_role_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


--
-- Name: user_role_id_seq1; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_role_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_role_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_role_id_seq1 OWNED BY public.user_role.id;


--
-- Name: user_role_module_list; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_role_module_list (
    id bigint NOT NULL,
    user_role_id bigint NOT NULL,
    user_module_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


--
-- Name: user_role_module_list_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_role_module_list_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_role_module_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_role_module_list_id_seq OWNED BY public.user_role_module_list.id;


--
-- Name: cart id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cart ALTER COLUMN id SET DEFAULT nextval('public.cart_id_seq1'::regclass);


--
-- Name: invoice id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invoice ALTER COLUMN id SET DEFAULT nextval('public.invoice_id_seq'::regclass);


--
-- Name: invoice_order_list id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invoice_order_list ALTER COLUMN id SET DEFAULT nextval('public.invoice_order_list_id_seq'::regclass);


--
-- Name: user_module id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_module ALTER COLUMN id SET DEFAULT nextval('public.user_module_id_seq1'::regclass);


--
-- Name: user_role id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_role ALTER COLUMN id SET DEFAULT nextval('public.user_role_id_seq1'::regclass);


--
-- Name: user_role_module_list id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_role_module_list ALTER COLUMN id SET DEFAULT nextval('public.user_role_module_list_id_seq'::regclass);


--
-- Data for Name: cart; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.cart (id, user_id, product_id, created_at, updated_at, deleted_at, quantity) FROM stdin;
\.


--
-- Data for Name: invoice; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.invoice (id, price, details, user_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: invoice_order_list; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.invoice_order_list (id, invoice_id, order_id) FROM stdin;
\.


--
-- Data for Name: order; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public."order" (id, buyer_id, seller_id, product_id, price, created_at, updated_at, deleted_at, quantity) FROM stdin;
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.product (id, name, description, photo, price, stock, slug, user_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public."user" (id, username, password, name, email, user_role_id, is_active, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: user_module; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_module (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	/admin/product/create	2025-07-03 02:19:45.004789	2025-07-03 02:19:45.004789	\N
2	/admin/product/update	2025-07-03 02:20:15.163835	2025-07-03 02:20:15.163835	\N
3	/admin/product/delete	2025-07-03 02:20:23.641968	2025-07-03 02:20:23.641968	\N
5	/cart/add-product	2025-07-03 04:39:23.448502	2025-07-03 04:39:23.448502	\N
4	/cart/	2025-07-03 04:39:18.112237	2025-07-03 04:39:18.112237	\N
6	/cart/buy	2025-07-03 05:18:12.893902	2025-07-03 05:18:12.893902	\N
7	/order/	2025-07-03 05:18:24.548542	2025-07-03 05:18:24.548542	\N
\.


--
-- Data for Name: user_role; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_role (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Merchant	2025-07-03 00:50:22.998435	2025-07-03 00:50:22.998435	\N
2	Customer	2025-07-03 00:50:22.998435	2025-07-03 00:50:22.998435	\N
\.


--
-- Data for Name: user_role_module_list; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_role_module_list (id, user_role_id, user_module_id, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	2025-07-03 02:41:44.193298	2025-07-03 02:41:44.193298	\N
2	1	2	2025-07-03 02:41:46.481891	2025-07-03 02:41:46.481891	\N
3	1	3	2025-07-03 02:41:48.327253	2025-07-03 02:41:48.327253	\N
4	2	4	2025-07-03 04:39:41.504892	2025-07-03 04:39:41.504892	\N
5	2	5	2025-07-03 04:39:44.129005	2025-07-03 04:39:44.129005	\N
6	2	6	2025-07-03 05:18:36.700769	2025-07-03 05:18:36.700769	\N
7	1	7	2025-07-03 05:18:39.599075	2025-07-03 05:18:39.599075	\N
\.


--
-- Name: cart_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.cart_id_seq', 1, false);


--
-- Name: cart_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.cart_id_seq1', 12, true);


--
-- Name: invoice_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.invoice_id_seq', 4, true);


--
-- Name: invoice_order_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.invoice_order_list_id_seq', 4, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.order_id_seq', 4, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.product_id_seq', 16, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_id_seq', 29, true);


--
-- Name: user_module_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_module_id_seq', 1, false);


--
-- Name: user_module_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_module_id_seq1', 7, true);


--
-- Name: user_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_role_id_seq', 2, true);


--
-- Name: user_role_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_role_id_seq1', 2, true);


--
-- Name: user_role_module_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_role_module_list_id_seq', 7, true);


--
-- Name: cart cart_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cart
    ADD CONSTRAINT cart_pkey PRIMARY KEY (id);


--
-- Name: invoice_order_list invoice_order_list_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invoice_order_list
    ADD CONSTRAINT invoice_order_list_pkey PRIMARY KEY (id);


--
-- Name: invoice invoice_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.invoice
    ADD CONSTRAINT invoice_pkey PRIMARY KEY (id);


--
-- Name: order order_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: product product_slug; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_slug UNIQUE (slug);


--
-- Name: user user_email; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email UNIQUE (email);


--
-- Name: user_module user_module_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_module
    ADD CONSTRAINT user_module_name UNIQUE (name);


--
-- Name: user_module user_module_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_module
    ADD CONSTRAINT user_module_pkey PRIMARY KEY (id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: user_role_module_list user_role_module_list_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_role_module_list
    ADD CONSTRAINT user_role_module_list_pkey PRIMARY KEY (id);


--
-- Name: user_role user_role_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_role
    ADD CONSTRAINT user_role_name UNIQUE (name);


--
-- Name: user_role user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_role
    ADD CONSTRAINT user_role_pkey PRIMARY KEY (id);


--
-- Name: user user_username; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_username UNIQUE (username);


--
-- Name: cart_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX cart_deleted_at ON public.cart USING btree (deleted_at);


--
-- Name: cart_product_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX cart_product_id ON public.cart USING btree (product_id);


--
-- Name: cart_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX cart_user_id ON public.cart USING btree (user_id);


--
-- Name: invoice_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX invoice_deleted_at ON public.invoice USING btree (deleted_at);


--
-- Name: invoice_order_list_invoice_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX invoice_order_list_invoice_id ON public.invoice_order_list USING btree (invoice_id);


--
-- Name: invoice_order_list_order_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX invoice_order_list_order_id ON public.invoice_order_list USING btree (order_id);


--
-- Name: invoice_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX invoice_user_id ON public.invoice USING btree (user_id);


--
-- Name: order_buyer_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX order_buyer_id ON public."order" USING btree (buyer_id);


--
-- Name: order_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX order_deleted_at ON public."order" USING btree (deleted_at);


--
-- Name: order_product_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX order_product_id ON public."order" USING btree (product_id);


--
-- Name: order_seller_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX order_seller_id ON public."order" USING btree (seller_id);


--
-- Name: product_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX product_deleted_at ON public.product USING btree (deleted_at);


--
-- Name: product_name; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX product_name ON public.product USING btree (name);


--
-- Name: product_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX product_user_id ON public.product USING btree (user_id);


--
-- Name: user_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_deleted_at ON public."user" USING btree (deleted_at);


--
-- Name: user_is_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_is_active ON public."user" USING btree (is_active);


--
-- Name: user_module_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_module_deleted_at ON public.user_module USING btree (deleted_at);


--
-- Name: user_role_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_role_deleted_at ON public.user_role USING btree (deleted_at);


--
-- Name: user_role_module_list_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_role_module_list_deleted_at ON public.user_role_module_list USING btree (deleted_at);


--
-- Name: user_role_module_list_user_module_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_role_module_list_user_module_id ON public.user_role_module_list USING btree (user_module_id);


--
-- Name: user_role_module_list_user_role_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_role_module_list_user_role_id ON public.user_role_module_list USING btree (user_role_id);


--
-- Name: user_user_role_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_user_role_id ON public."user" USING btree (user_role_id);


--
-- Name: cart cart_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cart
    ADD CONSTRAINT cart_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: cart cart_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cart
    ADD CONSTRAINT cart_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: order order_buyer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_buyer_id_fkey FOREIGN KEY (buyer_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: order order_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: order order_seller_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_seller_id_fkey FOREIGN KEY (seller_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: product product_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: user user_user_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_user_role_id_fkey FOREIGN KEY (user_role_id) REFERENCES public.user_role(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

