INSERT INTO advertTypes (id, advertType, dateCreation, isDeleted) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'House', NOW(), FALSE),
('6389b3d1-2bf3-43a0-bf1c-f55b51bc1a10', 'Flat', NOW(), FALSE),
('ebf69b2b-200b-4c4e-aa9b-47a7f56d6b94', 'House', NOW(), FALSE),
('fce05975-6c87-49a1-b59e-739ec2274c16', 'Flat', NOW(), FALSE),
('f3db14c1-4a47-4f5b-b155-42c46a3fe9ab', 'House', NOW(), FALSE),
('a2bf19b5-cc77-458e-8676-6286c89b9f01', 'Flat', NOW(), FALSE),
('97fb60b8-26e0-45c3-8202-1d0a1e8ee647', 'House', NOW(), FALSE),
('db0b0a75-8a45-4777-85f5-16e8c29b9353', 'Flat', NOW(), FALSE),
('9b4cde0e-894d-49f5-9e20-325a52a001fc', 'House', NOW(), FALSE),
('69f38a69-10da-45fb-ba0e-58929b88262e', 'Flat', NOW(), FALSE),
('a32a3eab-92d4-482f-b946-5b79d7d75a50', 'House', NOW(), FALSE),
('a6732f60-5d8f-4e2c-85b1-73c4d8f2b71b', 'Flat', NOW(), FALSE),
('e6c1d619-9c0c-4cf9-8c89-0c2e64ee56fb', 'House', NOW(), FALSE),
('ed95ed7f-34e7-44a2-9fb5-174b1e4d9ba8', 'Flat', NOW(), FALSE),
('731e760c-fbe1-4f3e-b113-8e2e2c3c5a36', 'House', NOW(), FALSE),
('ee452e26-0e5b-461b-b9e4-cf93b674b123', 'Flat', NOW(), FALSE),
('3f004961-956c-46a6-af6d-2fd71d66d5a5', 'House', NOW(), FALSE),
('b8d94f78-3bc1-4f82-a7f7-60310e759447', 'Flat', NOW(), FALSE),
('bd05f2e5-18ef-4e12-b9e3-9dc26b748d91', 'House', NOW(), FALSE),
('94b7b7c6-4f26-4a5b-869b-c7d8a5ecb123', 'Flat', NOW(), FALSE),
('ed78d4ee-fb1e-4646-bf7d-15502e8be133', 'House', NOW(), FALSE),
('8a9d3d59-4a2e-4f21-a473-d3de6cf40288', 'Flat', NOW(), FALSE),
('ae02583b-4070-4a2f-a631-3c956157d141', 'House', NOW(), FALSE),
('6c215f3a-39f5-4907-834c-25425e4a9a63', 'Flat', NOW(), FALSE),
('f3f723b8-87b7-40dc-96d0-6c1a65489a67', 'House', NOW(), FALSE),
('b7301b36-2c67-44a8-8772-4d36dcb4a71d', 'Flat', NOW(), FALSE),
('c3e78d50-6a20-4754-8009-433ca2f54c07', 'House', NOW(), FALSE),
('3b152e04-5a84-48db-9074-6f6c688529ee', 'Flat', NOW(), FALSE),
('089d0a75-f6a5-4512-86fd-55a3ed8e1748', 'House', NOW(), FALSE);

INSERT INTO users (id, passwordHash, levelUpdate, firstName, secondName, dateBirthday, phone, email, photo, dateCreation, isDeleted) VALUES
('18d5933e-b5e1-44f5-86d1-303e51891227', 'e38ad214943daad1d64c102faec29de4afe9da3d', 1, 'John', 'Doe', '1990-01-01', '1234567890', 'john.doe@example.com', 'photo1.jpg', NOW(), FALSE), -- password1 // SuperUserForComplexAdverts
('a4ff3cc1-cc31-4f14-80f0-7c6bf8c0b79c', '2aa60a8ff7fcd473d321e0146afd9e26df395147', 1, 'Jane', 'Smith', '1992-02-02', '2345678901', 'jane.smith@example.com', 'photo2.jpg', NOW(), FALSE), -- password2
('3cf39825-4521-4a86-8d9c-ec21a8cc1162', '1119cfd37ee247357e034a08d844eea25f6fd20f', 1, 'Alice', 'Johnson', '1994-03-03', '3456789012', 'alice.johnson@example.com', 'photo3.jpg', NOW(), FALSE), -- password3
('a37f2d87-2b20-4e2a-9d44-b71666e2a3e7', 'a1d7584daaca4738d499ad7082886b01117275d8', 1, 'Bob', 'Williams', '1996-04-04', '4567890123', 'bob.williams@example.com', 'photo4.jpg', NOW(), FALSE), -- password4
('4d76dbd3-47ff-497e-a6d6-828a6a3e5938', 'edba955d0ea15fdef4f61726ef97e5af507430c0', 1, 'Eve', 'Brown', '1998-05-05', '5678901234', 'eve.brown@example.com', 'photo5.jpg', NOW(), FALSE), -- password5
('5b76fd63-bc35-4a7f-b5ef-3c024ef79b97', '6d749e8a378a34cf19b4c02f7955f57fdba130a5', 1, 'Michael', 'Jones', '2000-06-06', '6789012345', 'michael.jones@example.com', 'photo6.jpg', NOW(), FALSE), -- password6
('e591697d-f2b1-4e35-8e26-ef4414c181e4', '330ba60e243186e9fa258f9992d8766ea6e88bc1', 1, 'Sarah', 'Lee', '2002-07-07', '7890123456', 'sarah.lee@example.com', 'photo7.jpg', NOW(), FALSE), -- password7
('b6e36a21-1b63-4e1d-bc94-6486339b91d7', 'a8dbbfa41cec833f8dd42be4d1fa9a13142c85c2', 1, 'Chris', 'Taylor', '2004-08-08', '8901234567', 'chris.taylor@example.com', 'photo8.jpg', NOW(), FALSE), -- password8
('72c7e5ae-d3e0-4e1a-ba68-6bc9a890ea91', '024b01916e3eaec66a2c4b6fc587b1705f1a6fc8', 1, 'Emma', 'Anderson', '2006-09-09', '9012345678', 'emma.anderson@example.com', 'photo9.jpg', NOW(), FALSE), -- password9
('ee9a1eb0-b968-4e2b-8ed4-583760e49e4d', 'f68ec41cde16f6b806d7b04c705766b7318fbb1d', 1, 'David', 'Martinez', '2008-10-10', '0123456789', 'david.martinez@example.com', 'photo10.jpg', NOW(), FALSE), -- password10
('8e7a9444-cd2e-47cb-89cb-ef17f1a79fc4', 'ddf6c9a1df4d57aef043ca8610a5a0dea097af0b', 1, 'Laura', 'Hernandez', '2010-11-11', '2234567890', 'laura.hernandez@example.com', 'photo11.jpg', NOW(), FALSE), -- password11
('f5dd6c3f-6656-4bc5-8e13-07a1a9cb4ed0', '10c28f9cf0668595d45c1090a7b4a2ae98edfa58', 1, 'Kevin', 'Young', '2012-12-12', '2345678902', 'kevin.young@example.com', 'photo12.jpg', NOW(), FALSE), -- password12
('3ff28968-6214-43cc-8987-fb9350c9140f', 'd505832286e2c1d2839f394de89b3af8dc3f8c1f', 1, 'Megan', 'Scott', '2014-01-13', '3456789022', 'megan.scott@example.com', 'photo13.jpg', NOW(), FALSE), -- password13
('570d3c8a-0db2-45eb-8f86-2a0d4ed2042a', '89f747bced37a9d8aee5c742e2aea373278eb29f', 1, 'Ryan', 'Nguyen', '2016-02-14', '4567890223', 'ryan.nguyen@example.com', 'photo14.jpg', NOW(), FALSE), -- password14
('3d40b9d2-9685-4f4d-aae8-3bc7e6a2dc7f', 'bd021e21c14628faa94d4aaac48c869d6b5d0ec3', 1, 'Katie', 'Kim', '2018-03-15', '5678902234', 'katie.kim@example.com', 'photo15.jpg', NOW(), FALSE), -- password15
('4cf846fc-3c75-4b4f-9668-8641eaa4bda8', '3de778e515e707114b622e769a308d1a2f84052b', 1, 'Brian', 'Singh', '2020-04-16', '6789022345', 'brian.singh@example.com', 'photo16.jpg', NOW(), FALSE), -- password16
('64b504d3-869b-470c-8f0d-0baed5c8e3de', 'b9c3d15c70a945d9e308ac763dd254b47c29bc0a', 1, 'Natalie', 'Garcia', '2022-05-17', '7890223456', 'natalie.garcia@example.com', 'photo17.jpg', NOW(), FALSE), -- password17
('3265834f-1cbb-4f35-9fbb-0ccf85b5113d', 'e7369527332f65fe86c44d87116801a0f4fbe5d3', 1, 'Alex', 'Perez', '2024-06-18', '8902234567', 'alex.perez@example.com', 'photo18.jpg', NOW(), FALSE), -- password18
('4763a74c-eb26-4e4d-9675-cb5c4d635035', '2c30de294b2ca17d5c356645a04ff4d0de832594', 1, 'Christine', 'Rodriguez', '2026-07-19', '9022345678', 'christine.rodriguez@example.com', 'photo19.jpg', NOW(), FALSE), -- password19
('f6c1e8e5-dc9f-4d8e-bf91-997f16fcf3c1', '6b00888703d6cae5654e2d2a6de79c42bbf94497', 1, 'Daniel', 'Lopez', '2028-08-20', '0223456789', 'daniel.lopez@example.com', 'photo20.jpg', NOW(), FALSE), -- password20
('764a774c-3dfb-442b-b7d5-0421616ff47c', '6bd1c0ac395c9cc40acd3fef59209944a8e09cd2', 1, 'Olivia', 'Hernandez', '2030-09-21', '3234567890', 'olivia.hernandez@example.com', 'photo21.jpg', NOW(), FALSE), -- password21
('e4c2f94d-3b84-4641-a78b-b832e70cf45b', '4393e23bbcfc18a7ff359b6130e73c55f5bdb541', 1, 'Matthew', 'Gonzalez', '2032-10-22', '2345678903', 'matthew.gonzalez@example.com', 'photo22.jpg', NOW(), FALSE), -- password22
('3f0b8fcf-6839-47f5-a1d0-d69ed1ba9b10', 'e5e080b98051e09a61175bdd4501701be7185582', 1, 'Ava', 'Wilson', '2034-11-23', '3456789032', 'ava.wilson@example.com', 'photo23.jpg', NOW(), FALSE), -- password23
('22c7682e-9f44-4d13-b0c1-bc5b53a317e3', 'ec962982c39b2137bc5453e66034a4e774164720', 1, 'Liam', 'Anderson', '2036-12-24', '4567890323', 'liam.anderson@example.com', 'photo24.jpg', NOW(), FALSE), -- password24
('f2b8d101-0212-47b3-8675-2fbc9d7ebf43', '493fa14b04d2bf8bb61eeaa9eca50bb1fbfc281d', 1, 'Emma', 'Martinez', '2038-01-25', '5678903234', 'emma.martinez@example.com', 'photo25.jpg', NOW(), FALSE), -- password25
('9302ef82-9ae1-41eb-b31a-96d364855b1e', '36d7048e5ff06f7707e4018ef1d17cf6c37dc0c5', 1, 'Noah', 'Brown', '2040-02-26', '6789032345', 'noah.brown@example.com', 'photo26.jpg', NOW(), FALSE), -- password26
('570a1c36-4be1-4586-920f-98f6b2b9c924', 'dff2565104b1a1e3b293579b35829abb47a73b2d', 1, 'Olivia', 'Nguyen', '2042-03-27', '7890323456', 'olivia.nguyen@example.com', 'photo27.jpg', NOW(), FALSE), -- password27
('7b8a872e-85c3-4abf-a3ac-57db872e5f92', 'c82940c8b3a430670709b2034b9423c728b34416', 1, 'William', 'Singh', '2044-04-28', '8903234567', 'william.singh@example.com', 'photo28.jpg', NOW(), FALSE), -- password28
('6ff0bc9a-354d-4bfb-ae5e-9e9e7c70d33c', '7cf05621019e9c633b84f2ba1ea097e8dae22bf7', 1, 'Sophia', 'Kim', '2046-05-29', '9032345678', 'sophia.kim@example.com', 'photo29.jpg', NOW(), FALSE), -- password29
('96b3efc8-8845-45cf-a8e7-1e3f46e155bc', '70de8f10edcbe18a77366264db2ee393c9827480', 1, 'James', 'Garcia', '2048-06-30', '0323456789', 'james.garcia@example.com', 'photo30.jpg', NOW(), FALSE); -- password30

INSERT INTO adverts (id, userId, advertTypeId, advertTypePlacement, title, description, phone, isAgent)
VALUES
('aafe24d6-eb4a-4e9f-82c5-7d3c92e3dc63', '18d5933e-b5e1-44f5-86d1-303e51891227', '550e8400-e29b-41d4-a716-446655440000', 'Sale', 'Beautiful House for Sale', 'Spacious and modern house with a garden.', '1234567890', FALSE),
('0b5a13eb-f7f5-45de-8a78-8d2e4fb9e433', 'a4ff3cc1-cc31-4f14-80f0-7c6bf8c0b79c', '6389b3d1-2bf3-43a0-bf1c-f55b51bc1a10', 'Rent', 'Cozy Flat for Rent', 'Fully furnished flat in a quiet neighborhood.', '2345678901', FALSE),
('bf394b49-2d64-45de-bbb1-5f9d4153542d', '3cf39825-4521-4a86-8d9c-ec21a8cc1162', 'ebf69b2b-200b-4c4e-aa9b-47a7f56d6b94', 'Sale', 'Spacious House with a View', 'Beautiful house overlooking the city.', '3456789012', FALSE),
('89d4e5c6-0f43-4047-bf8a-3ed4c2571d86', 'a37f2d87-2b20-4e2a-9d44-b71666e2a3e7', 'fce05975-6c87-49a1-b59e-739ec2274c16', 'Rent', 'Modern Flat in the City Center', 'Conveniently located flat with modern amenities.', '4567890123', FALSE),
('d69359a0-0c54-42d3-8c3a-2083c8c3e77b', '4d76dbd3-47ff-497e-a6d6-828a6a3e5938', 'f3db14c1-4a47-4f5b-b155-42c46a3fe9ab', 'Sale', 'Family House with a Garden', 'Cozy house with a spacious garden, perfect for families.', '5678901234', FALSE),
('1fc85c82-8ef3-4c61-983b-0630e5b29913', '5b76fd63-bc35-4a7f-b5ef-3c024ef79b97', 'a2bf19b5-cc77-458e-8676-6286c89b9f01', 'Rent', 'Studio Apartment for Rent', 'Cozy studio apartment in a vibrant neighborhood.', '6789012345', FALSE),
('eb60d0a7-f098-432f-8c42-f86cfa3d3019', 'e591697d-f2b1-4e35-8e26-ef4414c181e4', '97fb60b8-26e0-45c3-8202-1d0a1e8ee647', 'Sale', 'Luxury House with Pool', 'Stunning house with a swimming pool and a view.', '7890123456', FALSE),
('26e4898a-7f75-4dd1-997d-646f7f44417d', 'b6e36a21-1b63-4e1d-bc94-6486339b91d7', 'db0b0a75-8a45-4777-85f5-16e8c29b9353', 'Rent', 'Spacious Flat with Balcony', 'Bright flat with a balcony overlooking the park.', '8901234567', FALSE);

INSERT INTO buildings (id, floor, adress, adressPoint, yearCreation)
VALUES
('aafe24d6-eb4a-4e9f-82c5-7d3c92e3dc63', 2, '123 Main St', ST_GeographyFromText('POINT(-122.34900 47.65100)'), 2000),
('0b5a13eb-f7f5-45de-8a78-8d2e4fb9e433', 4, '456 Elm St', ST_GeographyFromText('POINT(-122.35000 47.65200)'), 1995),
('bf394b49-2d64-45de-bbb1-5f9d4153542d', 1, '789 Oak St', ST_GeographyFromText('POINT(-122.35100 47.65300)'), 2010),
('89d4e5c6-0f43-4047-bf8a-3ed4c2571d86', 3, '1010 Pine St', ST_GeographyFromText('POINT(-122.35200 47.65400)'), 2015),
('d69359a0-0c54-42d3-8c3a-2083c8c3e77b', 2, '1313 Maple St', ST_GeographyFromText('POINT(-122.35300 47.65500)'), 2005),
('1fc85c82-8ef3-4c61-983b-0630e5b29913', 5, '1515 Cedar St', ST_GeographyFromText('POINT(-122.35400 47.65600)'), 2018),
('eb60d0a7-f098-432f-8c42-f86cfa3d3019', 2, '1717 Walnut St', ST_GeographyFromText('POINT(-122.35500 47.65700)'), 2008),
('26e4898a-7f75-4dd1-997d-646f7f44417d', 4, '1919 Spruce St', ST_GeographyFromText('POINT(-122.35600 47.65800)'), 2012);

INSERT INTO flats (id, buildingId, advertTypeId, floor, ceilingHeight, squareGeneral, roomCount, squareResidential)
VALUES
('58c04b88-1f2e-4f17-a4ee-2b4792d67972', '0b5a13eb-f7f5-45de-8a78-8d2e4fb9e433', '6389b3d1-2bf3-43a0-bf1c-f55b51bc1a10', 4, 2.5, 100.0, 5, 80.0),
('3c4efb65-3d5e-4ec4-862f-62f62b224444', '89d4e5c6-0f43-4047-bf8a-3ed4c2571d86', 'fce05975-6c87-49a1-b59e-739ec2274c16', 3, 2.7, 120.0, 6, 100.0),
('8c9e50b8-91a0-485b-8a4f-cc4f71b5d555', '1fc85c82-8ef3-4c61-983b-0630e5b29913', 'a2bf19b5-cc77-458e-8676-6286c89b9f01', 5, 2.6, 80.0, 4, 60.0),
('40f5f022-6eb6-4c16-8ed5-66ef3b8b6666', '26e4898a-7f75-4dd1-997d-646f7f44417d', 'db0b0a75-8a45-4777-85f5-16e8c29b9353', 4, 2.8, 150.0, 3, 120.0);

INSERT INTO houses (id, buildingId, advertTypeId, ceilingHeight, squareArea, squareHouse, bedroomCount, statusArea, cottage, statusHome)
VALUES
('7b4f233e-68e1-459f-8a09-1094bb149e01', 'aafe24d6-eb4a-4e9f-82c5-7d3c92e3dc63', '550e8400-e29b-41d4-a716-446655440000', 3.0, 250.0, 200.0, 4, 'IHC', TRUE, 'Live'),
('80e45e7f-fc35-4ef0-b43f-b47b22c24046', 'bf394b49-2d64-45de-bbb1-5f9d4153542d', 'ebf69b2b-200b-4c4e-aa9b-47a7f56d6b94', 2.7, 180.0, 150.0, 3, 'DNP', FALSE, 'RepairNeed'),
('7c8631b2-0e0b-432f-8fd1-c1b24554f2a9', 'd69359a0-0c54-42d3-8c3a-2083c8c3e77b', 'f3db14c1-4a47-4f5b-b155-42c46a3fe9ab', 3.2, 220.0, 190.0, 4, 'G', TRUE, 'CompleteNeed'),
('b69b040c-d2b2-439d-87e0-f1044d440ae1', 'eb60d0a7-f098-432f-8c42-f86cfa3d3019', '97fb60b8-26e0-45c3-8202-1d0a1e8ee647', 3.5, 300.0, 250.0, 5, 'F', FALSE, 'Renovation');

WITH advertIds AS (
    SELECT id AS advertId, advertTypePlacement AS adType
    FROM adverts
)

INSERT INTO priceChanges (id, advertId, price)
SELECT
    uuid_generate_v4() AS id,
    advertIds.advertId,
    CASE 
        WHEN advertIds.adType = 'Sale' THEN
            ROUND((RANDOM() * (500012340 - 14000034) + 14430000)::numeric, 0)
        WHEN advertIds.adType = 'Rent' THEN
            ROUND((RANDOM() * (100000 - 54400) + 54300)::numeric, 0)
        ELSE
            0
    END AS price
FROM advertIds
JOIN adverts ON adverts.id = advertIds.advertId
JOIN advertTypes ON advertTypes.id = adverts.advertTypeId;

