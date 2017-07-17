DROP TABLE IF EXISTS fact;

CREATE TABLE fact
(
    id SERIAL PRIMARY KEY,
    factString TEXT,
    displayed BOOLEAN
);

INSERT INTO fact (factString, displayed) VALUES ('Odontophobia is the fear of teeth', true);
INSERT INTO fact (factString, displayed) VALUES ('Some penguins can leap 2-3 meters out of the water', true);
INSERT INTO fact (factString, displayed) VALUES ('A cockroach will live nine days without a head, before it starves to death', true);
INSERT INTO fact (factString, displayed) VALUES ('The thin line of cloud that forms behind an aircraft at high altitudes is called a contrail', true);
INSERT INTO fact (factString, displayed) VALUES ('In ancient Rome, it was considered a sign of leadership to be born with a crooked nose', true);
INSERT INTO fact (factString, displayed) VALUES ('The king of hearts is the only king without a moustache', true);
INSERT INTO fact (factString, displayed) VALUES ('The opening ceremony for the 2012 Olympics had 965 drummers there', true);
INSERT INTO fact (factString, displayed) VALUES ('The three most spoken english words are Hello, Stop and Taxi', true);
INSERT INTO fact (factString, displayed) VALUES ('On average, it takes 66 days to form a new habit', true);
INSERT INTO fact (factString, displayed) VALUES ('Snails can sleep for 3 years without eating', true);
INSERT INTO fact (factString, displayed) VALUES ('You are born with 300 bones; by the time you are an adult you will have 206', true);
INSERT INTO fact (factString, displayed) VALUES ('Life Savers werent meant to have a hole in them. it was a machine malfunction', true);
INSERT INTO fact (factString, displayed) VALUES ('Cats sleep 70% of their lives', true);
INSERT INTO fact (factString, displayed) VALUES ('The state of Florida is bigger than England', true);
INSERT INTO fact (factString, displayed) VALUES ('Experts at Intel say that micro processor speed will double every 18 months for at least the next 10 years', true);
INSERT INTO fact (factString, displayed) VALUES ('Karoke means "empty orchestra" in Japanese', true);
INSERT INTO fact (factString, displayed) VALUES ('Most lipstick contains fish scales', true);
INSERT INTO fact (factString, displayed) VALUES ('A rhinoceros horn is made of compacted hair', true);