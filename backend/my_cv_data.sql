--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'ru';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: education; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.education (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    duration character varying(255),
    college character varying(255),
    course character varying(255)
);


ALTER TABLE public.education OWNER TO postgres;

--
-- Name: education_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.education_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.education_id_seq OWNER TO postgres;

--
-- Name: education_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.education_id_seq OWNED BY public.education.id;


--
-- Name: experience; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.experience (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    duration character varying(255),
    title character varying(255),
    description text
);


ALTER TABLE public.experience OWNER TO postgres;

--
-- Name: experience_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.experience_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.experience_id_seq OWNER TO postgres;

--
-- Name: experience_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.experience_id_seq OWNED BY public.experience.id;


--
-- Name: skills; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.skills (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    type character varying(255),
    skill text[]
);


ALTER TABLE public.skills OWNER TO postgres;

--
-- Name: skills_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.skills_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.skills_id_seq OWNER TO postgres;

--
-- Name: skills_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.skills_id_seq OWNED BY public.skills.id;


--
-- Name: social; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.social (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    image character varying(255),
    link character varying(255)
);


ALTER TABLE public.social OWNER TO postgres;

--
-- Name: social_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.social_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.social_id_seq OWNER TO postgres;

--
-- Name: social_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.social_id_seq OWNED BY public.social.id;


--
-- Name: education id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.education ALTER COLUMN id SET DEFAULT nextval('public.education_id_seq'::regclass);


--
-- Name: experience id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experience ALTER COLUMN id SET DEFAULT nextval('public.experience_id_seq'::regclass);


--
-- Name: skills id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills ALTER COLUMN id SET DEFAULT nextval('public.skills_id_seq'::regclass);


--
-- Name: social id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.social ALTER COLUMN id SET DEFAULT nextval('public.social_id_seq'::regclass);


--
-- Data for Name: education; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.education (id, created_at, updated_at, deleted_at, duration, college, course) FROM stdin;
1	2025-11-09 18:25:32.169657	2025-11-09 18:25:32.169657	\N	2021-2023	Южноуральский государственный колледж	Разработчик веб и мультимедийных приложений
2	2025-11-09 19:44:26.145193	2025-11-09 19:53:29.728512	2025-11-09 20:03:08.520787	2023	Тест 2	Тест
3	2025-11-18 22:09:41.40142	2025-11-18 22:09:41.40142	\N	2024	Stepik	Курс «Программирование на Golang»
4	2025-11-18 22:13:26.467853	2025-11-18 22:13:26.467853	\N	2024	Karpov Courses	Курс «Симулятор SQL»
5	2025-11-18 22:14:12.808083	2025-11-18 22:14:12.808083	\N	2025	Stepik	Курс «Брокеры сообщений. Apache Kafka»
6	2025-11-18 22:14:52.030624	2025-11-18 22:14:52.030624	\N	2025	Stepik	Курс «Протокол HTTP простыми словами»
7	2025-11-18 22:15:46.53327	2025-11-18 22:15:46.53327	\N	2025	LA School	Курс «Работа с RabbitMQ в Go»
8	2025-11-18 22:16:16.646672	2025-11-18 22:16:16.646672	\N	2025	LA School	Курс «Основы работы с gRPC в Go»
9	2025-11-18 22:17:05.783797	2025-11-18 22:17:05.783797	\N	2025	Purple School	Курс «Продвинутый Golang»
10	2025-11-20 21:34:00.98409	2025-11-20 21:34:00.98409	2025-11-20 22:22:22.277082	test	test	test
11	2025-11-20 21:38:43.658176	2025-11-20 21:38:43.658176	2025-11-20 22:22:52.785294	test	test	test
12	2025-11-20 22:08:15.904483	2025-11-20 22:08:15.904483	2025-11-20 22:24:38.156275	test	test	test
\.


--
-- Data for Name: experience; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.experience (id, created_at, updated_at, deleted_at, duration, title, description) FROM stdin;
1	2025-11-09 21:13:33.576107	2025-11-19 19:41:11.740333	\N	Сентябрь 2025 - Настоящее время	Фриланс:  Fullstack-разработчик	<b>Backend-разработка (Go)</b><br><br>- Проектирование и разработка серверной части: полный цикл от продумывания архитектуры до реализации кодовой базы на Go.<br>- Реализация бизнес-логики: разработка REST API с использованием монолитной архитектуры.<br>- Работа с данными и аутентификацией: проектирование и подключение базы данных PostgreSQL; Использование ORM Gorm для построения запросов и взаимодействия с БД.<br>- Реализация механизмов безопасности: JWT-аутентификация с хешированием паролей через bcrypt.<br>- Интеграция и инфраструктура: настройка асинхронной обработки событий и сбора аналитики с использованием Apache Kafka;<br>- Развертывание изолированных тестовых сред (БД) с использованием Docker.<br>- Мониторинг: интеграция Grafana для визуализации метрик и мониторинга состояния сервисов.<br><br><b>Frontend-разработка</b><br><br>- Создание адаптивного пользовательского интерфейса с использованием HTML и фреймворка Tailwind CSS.<br>- Реализация клиентской логики на TypeScript: взаимодействие с backend через REST API (отправка данных и обработка ответов).<br><br><b>Сопутствующие обязанности</b><br><br>- Проектирование UI/UX: создание дизайна в Figma.<br>- Контроль версий и CI/CD: ведение проекта и организация рабочего процесса в GitLab.<br>- Внедрение SEO: расписывание meta-тегов для лучшей индексации поисковыми роботами.
2	2025-11-19 19:52:45.929373	2025-11-19 19:52:45.929373	\N	Октябрь 2024 — Август 2025	Сима-ленд: Golang-разработчик	<b>Разработка и оптимизация поисковых сервисов</b><br><br>- Реализация высокопроизводительного транслитератора для поисковой системы на Golang.<br>- Модификация алгоритмов поиска и поискового индексирования на базе Elasticsearch.<br>- Рефакторинг и расширение функциональности существующих сервисов.<br><br><b>Разработка бизнес-логики и API</b><br><br>- Полный цикл создания нового REST/gRPC эндпоинта для раздела "Купленные товары": от проектирования до интеграции.<br>- Генерация gRPC proto-файлов и реализация соответствующей серверной логики (хендлеров, репозиториев).<br><br><b>Работа с базами данных</b><br><br>- Разворачивание и настройка изолированной тестовой базы данных в Docker.<br><br><b>Тестирование и обеспечение качества</b><br><br>- Написание unit-тестов и использование mock-объектов для изоляции и проверки бизнес-логики.<br>- Интеграционное и функциональное тестирование API с помощью Postman.<br><br><b>Процессы разработки и инфраструктура</b><br><br>- Ведение разработки с использованием GitLab (контроль версий, code review, CI/CD).<br>- Активная работа с код-ревью: доработка кода по замечаниям тимлида, поддержание высокого стандарта кода.<br>- Сборка и управление Docker-образами через Docker Compose для тестовых сред.<br>- Интеграция новых модулей в существующую архитектуру приложения.
3	2025-11-19 19:57:35.402673	2025-11-19 19:57:35.402673	\N	Июль 2023 — Октябрь 2024	Сима-ленд: Менеджер проектов	<b>Управление проектами</b><br><br>- Полный цикл управления проектом: руководство на всех этапах жизненного цикла (от инициации и планирования до реализации, тестирования, запуска и сопровождения).<br>- Управление требованиями: сбор, анализ и формализация требований от заказчиков и стейкхолдеров, их приоритизация и перевод в конкретные задачи для команды разработки.<br>- Проектирование и документация: составление и согласование Технических Заданий (ТЗ) и другой проектной документации с использованием GitLab.<br>- Коммуникация и координация: организация эффективного взаимодействия между командами разработки, тестирования, заказчиками и другими стейкхолдерами. Выстраивание прозрачных рабочих процессов.<br>- Управление сроками: планирование, оценка трудозатрат и контроль сроков реализации проектных этапов.<br><br><b>Контроль качества и приемка работ</b><br><br>- Функциональное тестирование: проведение мануального тестирования функционала для проверки соответствия требованиям.<br>- Тестирование API: проверка доработок бэкенда через написание и выполнение тестовых сценариев в Postman.<br>- UI/UX: принятие решений по визуальным доработкам интерфейса, обеспечение соответствия продукта ожиданиям пользователя и утвержденным макетам.<br><br><b>Отчетная деятельность</b><br><br>- Подготовка и проведение отчетности: формирование и презентация межквартальных отчетов о статусе, прогрессе и результатах проектов для руководства.<br>- Ведение систем мониторинга: фиксация всех изменений проекта, обновлений статусов и ключевых метрик в локальной системе мониторинга (внутренняя разработка компании) для обеспечения актуальности информации и отслеживания истории проекта.
\.


--
-- Data for Name: skills; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.skills (id, created_at, updated_at, deleted_at, type, skill) FROM stdin;
3	2025-11-19 17:22:00.955293	2025-11-19 17:22:00.955293	\N	backend	{Golang,"REST API",HTTP,gRPC,Docker,Postman,"Apache Kafka",Redis,PostgreSQL,Grafana,Prometheus}
1	2025-11-19 17:09:03.721568	2025-11-19 17:09:03.721568	2025-11-19 17:23:47.197482	backend	\N
2	2025-11-19 17:15:39.110567	2025-11-19 17:15:39.110567	2025-11-19 17:24:20.369151	backend	\N
4	2025-11-19 18:32:53.800229	2025-11-19 18:32:53.800229	\N	Инструменты	{"Docker compose",Gitlab,Github,Postman,"IDE Goland",Figma}
5	2025-11-20 16:02:51.620932	2025-11-20 16:02:51.620932	\N	Frontend	{HTML/CSS,Swagger,OpenAPI,"Tailwind CSS",TypeScript}
6	2025-11-20 22:25:25.958794	2025-11-20 22:25:25.958794	2025-11-20 22:26:12.704757	test	{test}
\.


--
-- Data for Name: social; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.social (id, created_at, updated_at, deleted_at, image, link) FROM stdin;
3	2025-11-20 15:42:40.963393	2025-11-20 17:13:36.097044	\N	image/hhru.svg	https://ekaterinburg.hh.ru/resume/74f4e1a6ff0f91a2840039ed1f687875694238
2	2025-11-20 15:40:58.127418	2025-11-20 17:14:33.587482	\N	image/stack.svg	https://ru.stackoverflow.com/users/659291/maggot
1	2025-11-20 15:40:04.945277	2025-11-20 17:14:58.414369	\N	image/github.svg	https://github.com/mariapizzeria
\.


--
-- Name: education_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.education_id_seq', 12, true);


--
-- Name: experience_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.experience_id_seq', 3, true);


--
-- Name: skills_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.skills_id_seq', 6, true);


--
-- Name: social_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.social_id_seq', 3, true);


--
-- Name: education education_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.education
    ADD CONSTRAINT education_pkey PRIMARY KEY (id);


--
-- Name: experience experience_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experience
    ADD CONSTRAINT experience_pkey PRIMARY KEY (id);


--
-- Name: skills skills_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_pkey PRIMARY KEY (id);


--
-- Name: social social_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.social
    ADD CONSTRAINT social_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

