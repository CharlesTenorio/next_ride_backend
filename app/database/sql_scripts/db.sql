CREATE TABLE pedal_group (
	id_group varchar(250) NOT NULL,
	name_group varchar(50) NOT NULL,
	CONSTRAINT pedal_group_name_group_key UNIQUE (name_group),
	CONSTRAINT pedal_group_pkey PRIMARY KEY (id_group)
);


CREATE TABLE cyclist (
                         id_cyclist  varchar(250) PRIMARY KEY,
                         id_group varchar(250) NOT NULL,
                         cyclist_name  varchar ( 80 ) NOT NULL,
                         cpf varchar(15) unique not null,
                         birth date not null,
                         email varchar(150) unique not null,
                         blood_type varchar(10) not null,
                         health_plan varchar(80),
                         contact_emergency varchar(80) not null,
                         contact_fone varchar(20) not null,
                         got_to_know varchar(50) not null,
                         create_at  timestamp not null,
                         active bool not null,
                         img varchar(300),
                         participant_type varchar(40) not null,
                         pedaling int,
                         tours int,
                         travels int,
                         CONSTRAINT fk_group
                             FOREIGN KEY(id_group)
                                 REFERENCES pedal_group(id_group)

);
    






CREATE TABLE shedule (
	id_shedule serial PRIMARY KEY,
	id_group integer NOT NULL,
	day_of_event TIMESTAMP NOT NULL,
	day_of_week varchar ( 20 ) NOT NULL,
	event_type varchar ( 50 ) NOT NULL,
	maximun_cyclists integer NOT NULL,
	distance NUMERIC(9,2) NOT NULL,
	date_at TIMESTAMP NOT NULL,
	
	
	FOREIGN KEY (id_group) REFERENCES pedal_group (id_group)
);
	

	

CREATE TABLE chekin (
	id_chekin serial PRIMARY KEY,
	id_cyclist integer NOT NULL,
	id_group integer NOT NULL,
	id_shedule integer NOT NULL,
	date_at date NOT NULL,
	FOREIGN KEY (id_cyclist) REFERENCES cyclist (id_cyclist),
	FOREIGN KEY (id_group) REFERENCES pedal_group (id_group),
	FOREIGN KEY (id_shedule) REFERENCES shedule (id_shedule)
);

CREATE TABLE chekin_detail(
	id_chekin_detail serial PRIMARY KEY,
	id_chekin integer NOT NULL,
	id_group integer NOT NULL,
	id_cyclist integer NOT NULL,
	date_at date NOT NULL,
	time_at time NOT NULL,
	distance_at int NOT NULL,
	speed_at int NOT NULL,
	FOREIGN KEY (id_chekin) REFERENCES chekin (id_chekin),
	FOREIGN KEY (id_group) REFERENCES pedal_group (id_group),
	FOREIGN KEY (id_cyclist) REFERENCES cyclist (id_cyclist)
);
	


    