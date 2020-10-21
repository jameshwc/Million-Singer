INSERT INTO tours VALUES (1, "2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL);
INSERT INTO collects VALUES 
            (1,"2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "collect-1"),
            (2,"2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "collect-2"),
            (3,"2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "collect-3"),
            (4,"2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "collect-4"),
            (5,"2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "collect-5");
INSERT INTO songs VALUES
            (1, "2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "avLxcVkPgug", NULL, NULL, 
                "en", "Beautiful", "Eminem", "rap,hip-hop","15,23,29,33,37"),
            (2, "2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "JxzKNHfNRdI", NULL, NULL, 
                "en", "No Sleep", "Martin Garrix feat. Bonn", "edm","10,12,17"),
            (3, "2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "VDvr08sCPOc", NULL, NULL, 
                "en", "Remember The Name", "Fort Minor", "rap,hip-hop","7,11,14,22,24,28,31"),
            (4, "2020-10-19 22:39:15.769", "2020-10-19 22:39:15.769", NULL, "LqTfWEsGP4U", NULL, NULL, 
                "zh-tw", "那些勸我別抽菸的人都死了", "Fort Minor", "rap,hip-hop","7,11,14,22,24,28,31");
INSERT INTO tour_collects VALUES (1,1),(1,2),(1,3),(1,4),(1,5);
INSERT INTO collect_songs VALUES (1,1),(1,2),(1,3),(2,1),(2,2),(2,4),(3,1),(3,3),(3,4),(4,2),(4,3),(4,4),(5,1),(5,2),(5,3),(5,4);
INSERT INTO users VALUES (1,"2020-10-19 22:39:15.769","2020-10-19 22:39:15.769",NULL,"alice","alice@example.com","5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",0,NULL);