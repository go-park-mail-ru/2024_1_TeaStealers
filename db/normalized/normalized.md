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
        TIMESTAMP creation_date
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
        TIMESTAMP creation_date
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
        TIMESTAMP creation_date
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
        TIMESTAMP creation_date
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
        TIMESTAMP creation_date
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
        TIMESTAMP creation_date
        BOOLEAN is_deleted

    }
```
# Описание таблиц
## User
- level_update уровень jwt токена
- creation_date время создания аккаунта

## Advert
- user_id айди пользователа владельца объявления
- is_agent выставлено ли объявление риелтором
- priority какой рэйтинг продвижения объявления

## StatisticViewAdvert
### используется для просмотра объявлений конкретного пользователя

## FavouriteAdvert
### используется для хранения избранных объявлений пользователей

## Image
- photo содержит путь до фотографии
- priority приоритет изображения, определяет в каком порядке будут показываться в объявлении

## PriceChange
### содержит информацию об измениях цен в объявлениях
- price цена указанная в объявлении

## AdvertTypeHouse
### испоьзуется для загрузки объявлений с домами
- advert_id - id соотв. объявления
- house_id - id дома в объявлении

## AdvertTypeFlat
### аналогично как с AdvertTypeHouse

## House
### хранит информацию о параметрах дома
- building_id содержит информацию о материале и местоположении дома
- square_area площадь участка
- square_house площадь дома
- statis_area_house

## Flat
### хранит информацию о параметрах квартиры
- floor на каком этаже квартира
- square_general общая площадь квартиры
- bedroom_count однушка, двушка и т.д.
- square_residential - площадь спален?

## Building
### информация о здании
- material_building - из какого материала здание построено кирпич, блоки и т.п.
- address_point - координаты дома
- year_creation - год завершения постройки здания

## Complex
### информация о комплексе или доп информация о здании
- without_finishing_option - без отделки
- pre_finishing_option - с частичной отделкой
- finishing_option - с отделкой
- class_housing - статус комплекса
- security - охрана

## Company
### информация о застройщике
- photo - путь до логотипа компании