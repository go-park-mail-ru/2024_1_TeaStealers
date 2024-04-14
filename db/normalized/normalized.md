```mermaid
erDiagram
    User {
        BIGINT id PK
        TEXT password_hash
        INTEGER level_update
        TEXT first_name
        TEXT surname
        DATE date_birthdate
        TEXT phone
        TEXT email
        TEXT photo
        TIMESTAMP date_creation
        BOOLEAN is_deleted
    }

    FavouriteAdvert {
        BIGINT user_id PK,FK
        BIGINT advert_id PK,FK
        BOOLEAN is_deleted
    }

    FavouriteAdvert |{--|| User : user_favourite_adverts

    StatisticViewAdvert {
        BIGINT user_id PK,FK
        BIGINT advert_id PK,FK
        TIMESTAMP date_creation
        BOOLEAN is_deleted
    }

    StatisticViewAdvert ||--|| Advert : user_advert_link
    StatisticViewAdvert }|--|| User : user_adverts
    Advert ||--|| FavouriteAdvert : find_user_favourite_adverts


    AdvertTypeHouse {
        BIGINT house_id PK, FK
        BIGINT advert_id PK, FK
        BOOLEAN is_deleted
    }
    AdvertTypeHouse ||--|| Advert : get_advert_information
    AdvertTypeHouse ||--|| House : get_advert_information

    AdvertTypeFlat {
        BIGINT flat_id PK, FK
        BIGINT advert_id PK, FK
        BOOLEAN is_deleted
    }
    AdvertTypeFlat ||--}| Advert : get_advert_information
    AdvertTypeFlat }|--|| Flat : get_advert_information



    Flat {
        BIGINT id PK
        BIGINT building_id FK
        INTEGER floor
        FLOAT ceiling_height
        FLOAT square_general
        SMALLINT bedroom_count
        FLOAT square_residential
        BOOLEAN apartament
        BOOLEAN is_deleted
    }

    Flat }|--|| Building : get_flat_building_information

    House {
        BIGINT id PK
        BIGINT building_id FK
        FLOAT ceiling_height
        FLOAT square_area
        FLOAT square_house
        SMALLINT bedroom_count
        statusArea status_area_house
        BOOLEAN cottage
        statusHomeHouse status_home_house
        BOOLEAN is_deleted
    }
    House }|--|| Building : get_house_building_information

    Building {
        BIGINT id PK
        BIGINT complex_id FK
        material material_building
        TEXT address
        GEOGRAPHY address_point
        SMALLINT year_creation
        TIMESTAMP date_creation
        BOOLEAN is_deleted
    }
    Building ||--|| Complex : get_building_complex

    Complex {
        BIGINT id PK
        BIGINT company_id FK
        TEXT name
        TEXT address
        TEXT photo
        TEXT description
        DATE date_begin_build
        DATE date_end_build
        BOOLEAN without_finishing_option
        BOOLEAN finishing_option
        BOOLEAN pre_finishing_option
        classHouse class_housing
        BOOLEAN parking
        BOOLEAN security
        TIMESTAMP date_creation
        BOOLEAN is_deleted
    }
    Building ||--|| Complex : get_building_complex
    Complex ||--|| Company : get_complex_company



    Company {
        BIGINT id PK
        TEXT photo
        TEXT name
        SMALLINT creation_year
        TEXT phone
        TEXT description
        TIMESTAMP date_creation
        BOOLEAN is_deleted
    }
    Advert {
        BIGINT id PK
        BIGINT user_id FK
        TEXT title
        TEXT description
        TEXT phone
        BOOL is_agent
        SMALLINT priority
        TIMESTAMP date_creation
        BOOLEAN is_deleted
    }
    Advert ||--|{ Image : advert_images

    Image {
        BIGINT id PK
        BIGINT advert_id FK,UK
        TEXT photo
        SMALLINT priority UK
        TIMESTAMP date_creation
        BOOLEAN is_deleted

    }

    Advert ||--|{ PriceChange : advert_price_changes

    PriceChange {
        BIGINT id PK
        BIGINT advert_id FK
        NUMERIC price
        TIMESTAMP date_creation
        BOOLEAN is_deleted

    }
```
# Описание таблиц
## User
- id
- password_hash
- level_update
- first_name
- surname
- date_birthdate
- phone
- email
- photo
- date_creation
- is_deleted
- level_update - уровень jwt токена
- date_creation -  время создания аккаунта

{id} -> password_hash, level_update, first_name, surname, date_birthdate, phone, email, photo, date_creation, is_deleted

{phone} -> id, passwordHash, level_update, first_name, surname, date_birthdate, email, photo, dateCreation, isDeleted

{email} -> id, passwordHash, level_update, first_name, surname, date_birthdate, phone, photo, date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.


## Advert
- id
- user_id - айди пользователа владельца объявления
- title
- description
- phone
- is_agent - выставлено ли объявление риелтором
- priority - какой рэйтинг продвижения объявления
- date_creation
- is_deleted

{id} -> user_id, title, description, phone, priority, date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## StatisticViewAdvert
- user_id - для просмотра объявлений конкретного пользователя
- advert_id
- date_creation
- is_deleted

{user_id, advert_id} -> date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## FavouriteAdvert
- user_id - для просмотра избранных объявлений пользователем
- advert_id
- is_deleted

{user_id, advert_id} -> is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## Image
- id
- advert_id
- photo
- priority
- date_creation
- is_deleted
- photo - содержит путь до фотографии
- priority - приоритет изображения определяет в каком порядке будут показываться в объявлении

{id} -> advert_id, photo, priority, date_creation, is_deleted

{advert_id, priority} -> id, photo, date_creation,is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## PriceChange
- id
- advert_id
- price - цена указанная в объявлении
- date_creation
- is_deleted

{id} -> advert_id, price, date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## AdvertTypeHouse
- advert_id - id соотв. объявления
- house_id - id дома в объявлении
- is_deleted

{advert_id, house_id} -> is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## AdvertTypeFlat
- flat_id
- advert_id
- is_deleted

{advert_id, flat_id} -> is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## House
- id
- building_id содержит информацию о материале и местоположении дома
- ceiling_height
- square_area площадь участка
- square_house площадь дома
- bedroom_count
- status_area_house
- cottage
- status_home_house
- is_deleted

{id} -> building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## Flat
- id
- building_id
- floor - на каком этаже квартира
- ceiling_height
- square_general - общая площадь квартиры
- bedroom_count - однушка, двушка и т.д.
- square_residential - площадь спален
- apartament
- is_deleted

{id} -> building_id, floor, square_general, bedroom_count, square_residential, apartament,is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## Building
- id
- complex_id
- material_building - из какого материала здание построено кирпич, блоки и т.п.
- address
- address_point - координаты дома
- year_creation - год завершения постройки здания
- date_creation
- is_deleted

{id} -> complex_id, material_building, address, address_point, year_creation, date_creation, is_deleted

{address} -> id, complex_id, material_building, address_point, date_creation, is_deleted

{address_point} -> id, complex_id, material_building, address, date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## Complex
- id
- company_id
- name
- address
- photo
- description
- date_begin_build
- date_end_build
- without_finishing_option - без отделки
- pre_finishing_option - с частичной отделкой
- finishing_option - с отделкой
- class_housing - статус комплекса
- parking
- security - охрана
- date_creation
- is_deleted

{id} -> company_id, name, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option,
class_housing, parking, security, date_creation, is_deleted

{name} -> id, company_id, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option,
class_housing, parking, security, date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.

## Company
- id
- photo - путь до логотипа компании
- name
- creation_year
- phone
- description
- date_creation
- is_deleted

{id} -> photo, name, creation_year, phone, description, date_creation, is_deleted

Правилом первой формы является необходимость неделимости значения в каждом поле (столбце) строки – атомарность значений.
Первая форма - каждое поле в таблицах неделимое - атомарность значение
Вторая форма - нету зависимости неключевых полей от части составного ключа
Третья форма - нету зависимости неключевых полей от других неключевых полей
Переменная отношения находится в НФБК тогда и только тогда, когда для любой нетривиальной функциональной зависимости X→Y, X является надключом.