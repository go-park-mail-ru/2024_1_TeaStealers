package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"
)

const (
	InsertCreateAdvertTypeHouse  = `INSERT INTO advert_type_house (house_id, advert_id) VALUES ($1, $2)`
	InsertCreateAdvertTypeFlat   = `INSERT INTO advert_type_flat (flat_id, advert_id) VALUES ($1, $2)`
	InsertCreateAdvert           = `INSERT INTO advert (user_id, type_placement, title, description, phone, is_agent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	QueryProvinceId              = `SELECT id FROM province WHERE name=$1`
	QueryCreateProvince          = `INSERT INTO province (name) VALUES ($1) RETURNING id`
	QueryIdCreateTown            = `SELECT id FROM town WHERE name=$1 AND province_id=$2`
	InsertCreateTown             = `INSERT INTO town (name, province_id) VALUES ($1, $2) RETURNING id`
	QueryCreateStreet            = `SELECT id FROM street WHERE name=$1 AND town_id=$2`
	InsertCreateStreet           = `INSERT INTO street (name, town_id) VALUES ($1, $2) RETURNING id`
	QueryCreateHouseAddress      = `SELECT id FROM house_name WHERE name=$1 AND street_id=$2`
	InsertCreateHouseAddress     = `INSERT INTO house_name (name, street_id) VALUES ($1, $2) RETURNING id`
	QueryCreateAddress           = `SELECT id FROM address WHERE house_name_id=$1`
	InsertCreateAddress          = `INSERT INTO address (metro, house_name_id, address_point) VALUES ($1, $2, $3) RETURNING id`
	InsertCreatePriceChange      = `INSERT INTO price_change (advert_id, price) VALUES ($1, $2)`
	InsertCreateBuilding         = `INSERT INTO building (floor, material_building, address_id, year_creation) VALUES ($1, $2, $3, $4) RETURNING id`
	QueryCheckExistsBuilding     = `SELECT b.id, b.address_id, b.floor, b.material_building, b.year_creation FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`
	QueryCheckExistsBuildingData = `SELECT b.floor, b.material_building, b.year_creation, COALESCE(c.name, '') FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id LEFT JOIN complex AS c ON c.id=b.complex_id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`
	InsertCreateHouse            = `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	InsertCreateFlat             = `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	QuerySelectImages            = `SELECT id, photo, priority FROM image WHERE advert_id = $1 AND is_deleted = false`
	QuerySelectPriceChanges      = `SELECT price, created_at FROM price_change WHERE advert_id = $1 AND is_deleted = false`
	QueryGetTypeAdvertById       = `SELECT                   CASE
			WHEN ath.house_id IS NOT NULL THEN 'House'
			WHEN atf.flat_id IS NOT NULL THEN 'Flat'
			ELSE 'None'
		END AS type_advert FROM advert AS a LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id WHERE a.id=$1`
	QueryGetHouseAdvertById = `
		SELECT
	        a.id,
	        a.type_placement,
	        a.title,
	        a.description,
	        pc.price,
	        a.phone,
	        a.is_agent,
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
			ST_AsText(ad.address_point::geometry),
	        h.ceiling_height,
	        h.square_area,
	        h.square_house,
	        h.bedroom_count,
	        h.status_area_house,
	        h.cottage,
	        h.status_home_house,
	        b.floor,
	        b.year_creation,
	        COALESCE(b.material_building, 'Brick') as material,
	        a.created_at,
			CASE
				WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
				ELSE false
			END AS is_liked,
			CASE
				WHEN sva.advert_id IS NOT NULL THEN true
				ELSE false
			END AS is_viewed,
	        cx.id AS complexid,
	        c.photo AS companyphoto,
	        c.name AS companyname,
	        cx.name AS complexname
	    FROM
	        advert AS a
	    JOIN
	        advert_type_house AS at ON a.id = at.advert_id
	    JOIN
	        house AS h ON h.id = at.house_id
	    JOIN
	        building AS b ON h.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
		LEFT JOIN
			favourite_advert AS fa ON fa.advert_id=a.id AND fa.user_id=$2
		LEFT JOIN
			statistic_view_advert AS sva ON sva.advert_id=a.id AND sva.user_id=$2
	    LEFT JOIN
	        complex AS cx ON b.complex_id = cx.id
	    LEFT JOIN
	        company AS c ON cx.company_id = c.id
	    JOIN
	        LATERAL (
	            SELECT *
	            FROM price_change AS pc
	            WHERE pc.advert_id = a.id
	            ORDER BY pc.created_at DESC
	            LIMIT 1
	        ) AS pc ON TRUE
	    WHERE
	        a.id = $1 AND a.is_deleted = FALSE;`
	QueryCheckExistsFlat             = `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`
	QueryCheckExistsHouse            = `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`
	QueryGetIdTablesDeleteFlatAdvert = `
        SELECT
            f.id as flatid
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        WHERE a.id=$1;`
	QueryDeleteAdvertByIdDeleteFlatAdvert     = `UPDATE advert SET is_deleted=true WHERE id=$1;`
	QueryDeleteAdvertTypeByIdDeleteFlatAdvert = `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`
	QueryDeleteFlatByIdDeleteFlatAdvert       = `UPDATE flat SET is_deleted=true WHERE id=$1;`
	QueryDeletePriceChangesDeleteFlatAdvert   = `UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`
	QueryDeleteImagesDeleteFlatAdvert         = `UPDATE image SET is_deleted=true WHERE advert_id=$1;`
	QueryGetIdTablesDeleteHouseAdvert         = `
	        SELECT
	            h.id as houseid
	        FROM
	            advert AS a
	        JOIN
	            advert_type_house AS at ON a.id = at.advert_id
	        JOIN
	            house AS h ON h.id = at.house_id
	        WHERE a.id=$1;`
	QueryDeleteAdvertByIdDeleteHouseAdvert     = `UPDATE advert SET is_deleted=true WHERE id=$1;`
	QueryDeleteAdvertTypeByIdDeleteHouseAdvert = `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
	QueryDeleteHouseByIdDeleteHouseAdvert      = `UPDATE house SET is_deleted=true WHERE id=$1;`
	QueryDeletePriceChangesDeleteHouseAdvert   = `UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`
	QueryDeleteImagesDeleteHouseAdvert         = `UPDATE image SET is_deleted=true WHERE advert_id=$1;`

	QueryChangeTypeAdvert = `SELECT 			CASE
					WHEN ath.house_id IS NOT NULL THEN 'House'
					WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					ELSE 'None'
				END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`

	QuerySelectBuildingIdByFlatChangeTypeAdvert = `SELECT b.id AS buildingid, f.id AS flatid  FROM advert AS a JOIN advert_type_flat AS at ON at.advert_id=a.id JOIN flat AS f ON f.id=at.flat_id JOIN building AS b ON f.building_id=b.id WHERE a.id=$1`

	QuerySelectBuildingIdByHouseChangeTypeAdvert = `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`

	QueryInsertFlatChangeTypeAdvert = `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament)
					VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	QueryInsertHouseChangeTypeAdvert = `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	QueryInsertTypeFlatChangeTypeAdvert = `INSERT INTO advert_type_flat (advert_id, flat_id) VALUES ($1, $2);`

	QueryInsertTypeHouseChangeTypeAdvert = `INSERT INTO advert_type_house (advert_id, house_id) VALUES ($1, $2);`

	QueryRestoreFlatByIdChangeTypeAdvert = `UPDATE flat SET is_deleted=false WHERE id=$1;`

	QueryRestoreHouseByIdChangeTypeAdvert = `UPDATE house SET is_deleted=false WHERE id=$1;`

	QueryDeleteFlatByIdChangeTypeAdvert = `UPDATE flat SET is_deleted=true WHERE id=$1;`

	QueryDeleteHouseByIdChangeTypeAdvert = `UPDATE house SET is_deleted=true WHERE id=$1;`

	QueryDeleteAdvertTypeFlatChangeTypeAdvert = `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`

	QueryDeleteAdvertTypeHouseChangeTypeAdvert = `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`

	QueryRestoreAdvertTypeFlatChangeTypeAdvert = `UPDATE advert_type_flat SET is_deleted=false WHERE advert_id=$1 AND flat_id=$2;`

	QueryRestoreAdvertTypeHouseChangeTypeAdvert = `UPDATE advert_type_house SET is_deleted=false WHERE advert_id=$1 AND house_id=$2;`

	QueryGetIdTables = `
		        SELECT
		            b.id as buildingid,
		            h.id as houseid,
		            pc.price
		        FROM
		            advert AS a
		        JOIN
		            advert_type_house AS at ON a.id = at.advert_id
		        JOIN
		            house AS h ON h.id = at.house_id
		        JOIN
		            building AS b ON h.building_id = b.id
		        LEFT JOIN
		            LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		        WHERE a.id=$1;`

	QueryUpdateAdvertByIdUpdateHouseAdvert   = `UPDATE advert SET type_placement=$1, title=$2, description=$3, phone=$4, is_agent=$5 WHERE id=$6;`
	QueryUpdateBuildingByIdUpdateHouseAdvert = `UPDATE building SET floor=$1, material_building=$2, address_id=$3, year_creation=$4 WHERE id=$5;`
	QueryUpdateHouseByIdUpdateHouseAdvert    = `UPDATE house SET ceiling_height=$1, square_area=$2, square_house=$3, bedroom_count=$4, status_area_house=$5, cottage=$6, status_home_house=$7 WHERE id=$8;`
	QueryInsertPriceChangeUpdateHouseAdvert  = `INSERT INTO price_change (advert_id, price)
            VALUES ($1, $2)`

	QueryGetIdTablesUpdateFlatAdvert = `
                    SELECT
                        b.id as buildingid,
                        f.id as flatid,
                        pc.price
                    FROM
                        advert AS a
                    JOIN
                        advert_type_flat AS at ON a.id = at.advert_id
                    JOIN
                        flat AS f ON f.id = at.flat_id
                    JOIN
                        building AS b ON f.building_id = b.id
                    LEFT JOIN
                        LATERAL (
                            SELECT *
                            FROM price_change AS pc
                            WHERE pc.advert_id = a.id
                            ORDER BY pc.created_at DESC
                            LIMIT 1
                        ) AS pc ON TRUE
                    WHERE a.id=$1;`

	QueryUpdateAdvertByIdUpdateFlatAdvert   = `UPDATE advert SET type_placement=$1, title=$2, description=$3, phone=$4, is_agent=$5 WHERE id=$6;`
	QueryUpdateBuildingByIdUpdateFlatAdvert = `UPDATE building SET floor=$1, material_building=$2, address_id=$3, year_creation=$4 WHERE id=$5;`
	QueryUpdateFlatByIdUpdateFlatAdvert     = `UPDATE flat SET floor=$1, ceiling_height=$2, square_general=$3, bedroom_count=$4, square_residential=$5, apartament=$6 WHERE id=$7;`

	QueryInsertPriceChangeUpdateFlatAdvert = `INSERT INTO price_change (advert_id, price)
            VALUES ($1, $2)`
	QueryGetIdTablesUpdateHouseAdvert = `
		        SELECT
		            b.id as buildingid,
		            h.id as houseid,
		            pc.price
		        FROM
		            advert AS a
		        JOIN
		            advert_type_house AS at ON a.id = at.advert_id
		        JOIN
		            house AS h ON h.id = at.house_id
		        JOIN
		            building AS b ON h.building_id = b.id
		        LEFT JOIN
		            LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		        WHERE a.id=$1;`
	QueryGetFlatAdvertById = `
            	SELECT
            		a.id,
            		a.type_placement,
            		a.title,
            		a.description,
            		pc.price,
            		a.phone,
            		a.is_agent,
            		ad.metro,
            		hn.name,
            		s.name,
            		t.name,
            		p.name,
            		ST_AsText(ad.address_point::geometry),
                    f.floor,
                    f.ceiling_height,
                    f.square_general,
                    f.bedroom_count,
                    f.square_residential,
                    f.apartament,
                    b.floor AS floorGeneral,
                    b.year_creation,
                    COALESCE(b.material_building, 'Brick') as material,
                    a.created_at,
            		CASE
            			WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
            			ELSE false
            		END AS is_liked,
            		CASE
            			WHEN sva.advert_id IS NOT NULL THEN true
            			ELSE false
            		END AS is_viewed,
                    cx.id AS complexid,
                    c.photo AS companyphoto,
                    c.name AS companyname,
                    cx.name AS complexname
                FROM
                    advert AS a
                JOIN
                    advert_type_flat AS at ON a.id = at.advert_id
                JOIN
                    flat AS f ON f.id = at.flat_id
                JOIN
                    building AS b ON f.building_id = b.id
            		JOIN address AS ad ON b.address_id=ad.id
            		JOIN house_name AS hn ON hn.id=ad.house_name_id
            		JOIN street AS s ON s.id=hn.street_id
            		JOIN town AS t ON t.id=s.town_id
            		JOIN province AS p ON p.id=t.province_id
            	LEFT JOIN
            		favourite_advert AS fa ON fa.advert_id=a.id AND fa.user_id=$2
            	LEFT JOIN
            		statistic_view_advert AS sva ON sva.advert_id=a.id AND sva.user_id=$2
                LEFT JOIN
                    complex AS cx ON b.complex_id = cx.id
                LEFT JOIN
                    company AS c ON cx.company_id = c.id
                LEFT JOIN
                    LATERAL (
                        SELECT *
                        FROM price_change AS pc
                        WHERE pc.advert_id = a.id
                        ORDER BY pc.created_at DESC
                        LIMIT 1
                    ) AS pc ON TRUE
                WHERE
                    a.id = $1 AND a.is_deleted = FALSE;`

	QueryBaseAdvertGetSquareAdverts = `
         SELECT
             a.id,
             a.type_placement,
 			CASE
            		WHEN ath.house_id IS NOT NULL THEN 'House'
            		WHEN atf.flat_id IS NOT NULL THEN 'Flat'
            		ELSE 'None'
        		END AS type_advert,
             i.photo,
             pc.price,
             a.created_at
         FROM
             advert AS a
 			LEFT JOIN advert_type_house AS ath ON ath.advert_id=a.id
 			LEFT JOIN advert_type_flat AS atf ON atf.advert_id=a.id
             LEFT JOIN LATERAL (
                 SELECT *
                 FROM price_change AS pc
                 WHERE pc.advert_id = a.id
                 ORDER BY pc.created_at DESC
                 LIMIT 1
             ) AS pc ON TRUE
             JOIN image AS i ON i.advert_id = a.id
         WHERE i.priority = (
                 SELECT MIN(priority)
                 FROM image
                 WHERE advert_id = a.id
                     AND is_deleted = FALSE
             )
             AND i.is_deleted = FALSE
         ORDER BY
             a.created_at DESC
         LIMIT $1
         OFFSET $2;`

	QueryFlatGetSquareAdverts = `
        SELECT
            f.square_general,
            f.floor,
            ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            b.floor AS floorgeneral,
            f.bedroom_count
        FROM
            advert AS a
            JOIN advert_type_flat AS at ON a.id = at.advert_id
            JOIN flat AS f ON f.id=at.flat_id
            JOIN building AS b ON f.building_id=b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id=$1 AND a.is_deleted = FALSE
        ORDER BY
            a.created_at DESC;`
	QueryHouseGetSquareAdverts = `
        	SELECT
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            h.cottage,
            h.square_house,
            h.square_area,
            h.bedroom_count,
            b.floor
        FROM
            advert AS a
            JOIN advert_type_house AS at ON a.id = at.advert_id
            JOIN house AS h ON h.id=at.house_id
            JOIN building AS b ON h.building_id=b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id=$1
        ORDER BY
            a.created_at DESC;`

	QueryFlatGetRectangleAdverts = `
		        SELECT
		            f.square_general,
		            f.floor,
					ST_AsText(ad.address_point::geometry),
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id = at.flat_id
		            JOIN building AS b ON f.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
	QueryHouseGetRectangleAdverts = `
		        SELECT
					ST_AsText(ad.address_point::geometry),
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            b.floor
		        FROM
					advert AS a
					JOIN advert_type_house AS at ON a.id = at.advert_id
					JOIN house AS h ON h.id = at.house_id
					JOIN building AS b ON h.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`

	QueryBaseAdvertGetRectangleAdvertsByUser = `
	        SELECT
				a.id,
				a.title,
				a.description,
				CASE
				   WHEN ath.house_id IS NOT NULL THEN 'House'
				   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
				   ELSE 'None'
			    END AS type_advert,
	            CASE
	                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
	                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
	                ELSE 0
	            END AS rcount,
	            a.phone,
	            a.type_placement,
	            pc.price,
	            i.photo,
	            a.created_at,
				CASE
					WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
					ELSE false
				END AS is_liked
	        FROM
	            advert AS a
	            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
				LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
				LEFT JOIN favourite_advert AS fa ON a.id=fa.advert_id AND fa.user_id=$1
	            LEFT JOIN flat AS f ON f.id = atf.flat_id
	            LEFT JOIN house AS h ON h.id = ath.house_id
	            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
	            LEFT JOIN LATERAL (
	                SELECT *
	                FROM price_change AS pc
	                WHERE pc.advert_id = a.id
	                ORDER BY pc.created_at DESC
	                LIMIT 1
	            ) AS pc ON TRUE
	            JOIN image AS i ON i.advert_id = a.id
	        WHERE i.priority = (
	                SELECT MIN(priority)
	                FROM image
	                WHERE advert_id = a.id
	                    AND is_deleted = FALSE
	            )
	            AND i.is_deleted = FALSE
	            AND a.is_deleted = FALSE
	            AND a.user_id = $1
				ORDER BY a.created_at DESC
				LIMIT $2
				OFFSET $3`
	QueryFlatGetRectangleAdvertsByUser = `
	        SELECT
	            f.square_general,
	            f.floor,
				ad.metro,
				hn.name,
				s.name,
				t.name,
				p.name,
	            b.floor AS floorgeneral
	        FROM
	            advert AS a
	            JOIN advert_type_flat AS at ON a.id = at.advert_id
	            JOIN flat AS f ON f.id = at.flat_id
	            JOIN building AS b ON f.building_id = b.id
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
	        WHERE a.id = $1
	        ORDER BY
	            a.created_at DESC;`
	QueryHouseGetRectangleAdvertsByUser = `
	        SELECT
				ad.metro,
				hn.name,
				s.name,
				t.name,
				p.name,
	            h.cottage,
	            h.square_house,
	            h.square_area,
	            b.floor
	        FROM
				advert AS a
				JOIN advert_type_house AS at ON a.id = at.advert_id
				JOIN house AS h ON h.id = at.house_id
				JOIN building AS b ON h.building_id = b.id
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
	        WHERE a.id = $1
	        ORDER BY
	            a.created_at DESC;`

	QueryBaseAdvertGetRectangleAdvertsByComplex = `
		        SELECT
					a.id,
					a.title,
					a.description,
					CASE
					   WHEN ath.house_id IS NOT NULL THEN 'House'
					   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					   ELSE 'None'
				    END AS type_advert,
		            CASE
		                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
		                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
		                ELSE 0
		            END AS rcount,
		            a.phone,
		            a.type_placement,
		            pc.price,
		            i.photo,
		            a.created_at
		        FROM
		            advert AS a
		            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
					LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
		            LEFT JOIN flat AS f ON f.id = atf.flat_id
		            LEFT JOIN house AS h ON h.id = ath.house_id
		            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		            AND a.is_deleted = FALSE
		            AND b.complex_id = $1
					ORDER BY a.created_at DESC
					LIMIT $2
					OFFSET $3`
	QueryFlatGetRectangleAdvertsByComplex = `
		        SELECT
		            f.square_general,
		            f.floor,
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id = at.flat_id
		            JOIN building AS b ON f.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
	QueryHouseGetRectangleAdvertsByComplex = `
		        SELECT
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            b.floor
		        FROM
					advert AS a
					JOIN advert_type_house AS at ON a.id = at.advert_id
					JOIN house AS h ON h.id = at.house_id
					JOIN building AS b ON h.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`

	QueryBaseAdvertGetRectangleAdvertsLikedByUser = `
	        SELECT
				a.id,
				a.title,
				a.description,
				CASE
				   WHEN ath.house_id IS NOT NULL THEN 'House'
				   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
				   ELSE 'None'
			    END AS type_advert,
	            CASE
	                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
	                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
	                ELSE 0
	            END AS rcount,
	            a.phone,
	            a.type_placement,
	            pc.price,
	            i.photo,
	            a.created_at,
				CASE
					WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
					ELSE false
				END AS is_liked
	        FROM
	            advert AS a
	            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
				LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
				JOIN favourite_advert AS fa ON (a.id=fa.advert_id AND fa.user_id=$1 AND fa.is_deleted = false)
	            LEFT JOIN flat AS f ON f.id = atf.flat_id
	            LEFT JOIN house AS h ON h.id = ath.house_id
	            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
	            LEFT JOIN LATERAL (
	                SELECT *
	                FROM price_change AS pc
	                WHERE pc.advert_id = a.id
	                ORDER BY pc.created_at DESC
	                LIMIT 1
	            ) AS pc ON TRUE
	            JOIN image AS i ON i.advert_id = a.id
	        WHERE i.priority = (
	                SELECT MIN(priority)
	                FROM image
	                WHERE advert_id = a.id
	                    AND is_deleted = FALSE
	            )
	            AND i.is_deleted = FALSE
	            AND a.is_deleted = FALSE
	            AND a.user_id = $1
				ORDER BY a.created_at DESC
				LIMIT $2
				OFFSET $3`
	QueryFlatGetRectangleAdvertsLikedByUser = `
	        SELECT
	            f.square_general,
	            f.floor,
				ad.metro,
				hn.name,
				s.name,
				t.name,
				p.name,
	            b.floor AS floorgeneral
	        FROM
	            advert AS a
	            JOIN advert_type_flat AS at ON a.id = at.advert_id
	            JOIN flat AS f ON f.id = at.flat_id
	            JOIN building AS b ON f.building_id = b.id
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
	        WHERE a.id = $1
	        ORDER BY
	            a.created_at DESC;`
	QueryHouseGetRectangleAdvertsLikedByUser = `
	        SELECT
				ad.metro,
				hn.name,
				s.name,
				t.name,
				p.name,
	            h.cottage,
	            h.square_house,
	            h.square_area,
	            b.floor
	        FROM
				advert AS a
				JOIN advert_type_house AS at ON a.id = at.advert_id
				JOIN house AS h ON h.id = at.house_id
				JOIN building AS b ON h.building_id = b.id
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
	        WHERE a.id = $1
	        ORDER BY
	            a.created_at DESC;`
)

// AdvertRepo represents a repository for adverts changes.
type AdvertRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of AdvertRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *AdvertRepo {
	return &AdvertRepo{db: db, logger: logger}
}

func (r *AdvertRepo) BeginTx(ctx context.Context) (models.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.BeginTxMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.BeginTxMethod)
	return tx, nil
}

// CreateAdvertTypeHouse creates a new advertTypeHouse in the database.
func (r *AdvertRepo) CreateAdvertTypeHouse(ctx context.Context, tx models.Transaction, newAdvertType *models.HouseTypeAdvert) error {
	// insert := `INSERT INTO advert_type_house (house_id, advert_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, InsertCreateAdvertTypeHouse, newAdvertType.HouseID, newAdvertType.AdvertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvertTypeFlat creates a new advertTypeFlat in the database.
func (r *AdvertRepo) CreateAdvertTypeFlat(ctx context.Context, tx models.Transaction, newAdvertType *models.FlatTypeAdvert) error {
	// insert := `INSERT INTO advert_type_flat (flat_id, advert_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, InsertCreateAdvertTypeFlat, newAdvertType.FlatID, newAdvertType.AdvertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *AdvertRepo) CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) (int64, error) {
	// insert := `INSERT INTO advert (user_id, type_placement, title, description, phone, is_agent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var idAdvert int64
	if err := tx.QueryRowContext(ctx, InsertCreateAdvert, newAdvert.UserID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority).Scan(&idAdvert); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return idAdvert, nil
}

// CreateProvince creates a new province in the database.
func (r *AdvertRepo) CreateProvince(ctx context.Context, tx models.Transaction, name string) (int64, error) {
	// QueryProvinceId = `SELECT id FROM province WHERE name=$1`

	res := r.db.QueryRow(QueryProvinceId, name)

	var provinceId int64
	if err := res.Scan(&provinceId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return provinceId, nil
	}

	// insert := `INSERT INTO province (name) VALUES ($1) RETURNING id`
	if err := tx.QueryRowContext(ctx, QueryCreateProvince, name).Scan(&provinceId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return provinceId, nil
}

// CreateTown creates a new town in the database.
func (r *AdvertRepo) CreateTown(ctx context.Context, tx models.Transaction, idProvince int64, name string) (int64, error) {
	// QueryIdTownProvince = `SELECT id FROM town WHERE name=$1 AND province_id=$2`

	res := r.db.QueryRow(QueryIdCreateTown, name, idProvince)

	var townId int64
	if err := res.Scan(&townId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return townId, nil
	}

	// InsertCreateTown = `INSERT INTO town (name, province_id) VALUES ($1, $2) RETURNING id`
	if err := tx.QueryRowContext(ctx, InsertCreateTown, name, idProvince).Scan(&townId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return townId, nil
}

// CreateStreet creates a new street in the database.
func (r *AdvertRepo) CreateStreet(ctx context.Context, tx models.Transaction, idTown int64, name string) (int64, error) {
	// QueryCreateStreet = `SELECT id FROM street WHERE name=$1 AND town_id=$2`

	res := r.db.QueryRow(QueryCreateStreet, name, idTown)

	var streetId int64
	if err := res.Scan(&streetId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return streetId, nil
	}

	// InsertCreateStreet = `INSERT INTO street (name, town_id) VALUES ($1, $2) RETURNING id`
	if err := tx.QueryRowContext(ctx, InsertCreateStreet, name, idTown).Scan(&streetId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return streetId, nil
}

// CreateHouseAddress creates a new house in the database.
func (r *AdvertRepo) CreateHouseAddress(ctx context.Context, tx models.Transaction, idStreet int64, name string) (int64, error) {
	// QueryCreateHouseAddress = `SELECT id FROM house_name WHERE name=$1 AND street_id=$2`

	res := r.db.QueryRow(QueryCreateHouseAddress, name, idStreet)

	var houseId int64
	if err := res.Scan(&houseId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return houseId, nil
	}

	// InsertCreateHouseAddress = `INSERT INTO house_name (name, street_id) VALUES ($1, $2) RETURNING id`
	if err := tx.QueryRowContext(ctx, InsertCreateHouseAddress, name, idStreet).Scan(&houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return houseId, nil
}

// CreateAddress creates a new address in the database.
func (r *AdvertRepo) CreateAddress(ctx context.Context, tx models.Transaction, idHouse int64, metro string, address_point string) (int64, error) {
	// QueryCreateAddress = `SELECT id FROM address WHERE house_name_id=$1`

	res := r.db.QueryRow(QueryCreateAddress, idHouse)

	var addressId int64
	if err := res.Scan(&addressId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return addressId, nil
	}

	// InsertCreateAddress = `INSERT INTO address (metro, house_name_id, address_point) VALUES ($1, $2, $3) RETURNING id`
	if err := tx.QueryRowContext(ctx, InsertCreateAddress, metro, idHouse, address_point).Scan(&addressId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return addressId, nil
}

// CreatePriceChange creates a new price change in the database.
func (r *AdvertRepo) CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error {
	// InsertCreatePriceChange = `INSERT INTO price_change (advert_id, price) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, InsertCreatePriceChange, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreatePriceChangeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreatePriceChangeMethod)
	return nil
}

// CreateBuilding creates a new building in the database.
func (r *AdvertRepo) CreateBuilding(ctx context.Context, tx models.Transaction, newBuilding *models.Building) (int64, error) {
	// InsertCreateBuilding = `INSERT INTO building (floor, material_building, address_id, year_creation) VALUES ($1, $2, $3, $4) RETURNING id`
	var idBuilding int64
	if err := tx.QueryRowContext(ctx, InsertCreateBuilding, newBuilding.Floor, newBuilding.Material, newBuilding.AddressID, newBuilding.YearCreation).Scan(&idBuilding); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateBuildingMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateBuildingMethod)
	return idBuilding, nil
}

// CheckExistsBuilding check exists building.
func (r *AdvertRepo) CheckExistsBuilding(ctx context.Context, adress *models.AddressData) (*models.Building, error) {
	// QueryCheckExistsBuilding = `SELECT b.id, b.address_id, b.floor, b.material_building, b.year_creation FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`

	building := &models.Building{}

	res := r.db.QueryRowContext(ctx, QueryCheckExistsBuilding, adress.Province, adress.Town, adress.Street, adress.House)

	if err := res.Scan(&building.ID, &building.AddressID, &building.Floor, &building.Material, &building.YearCreation); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod)
	return building, nil
}

// CheckExistsBuildingData check exists buildings. Нужна для выпадающего списка существующих зданий по адресу(Для создания объявления)
func (r *AdvertRepo) CheckExistsBuildingData(ctx context.Context, adress *models.AddressData) (*models.BuildingData, error) {
	// QueryCheckExistsBuildingData = `SELECT b.floor, b.material_building, b.year_creation, COALESCE(c.name, '') FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id LEFT JOIN complex AS c ON c.id=b.complex_id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`

	building := &models.BuildingData{}

	res := r.db.QueryRowContext(ctx, QueryCheckExistsBuildingData, adress.Province, adress.Town, adress.Street, adress.House)

	if err := res.Scan(&building.Floor, &building.Material, &building.YearCreation, &building.ComplexName); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod, err)
		return nil, nil
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod)
	return building, nil
}

// CreateHouse creates a new house in the database.
func (r *AdvertRepo) CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) (int64, error) {
	// InsertCreateHouse = `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var lastInsertID int64
	if err := tx.QueryRowContext(ctx, InsertCreateHouse, newHouse.BuildingID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome).Scan(&lastInsertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateHouseMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateHouseMethod)
	return lastInsertID, nil
}

// CreateFlat creates a new flat in the database.
func (r *AdvertRepo) CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) (int64, error) {
	// InsertCreateFlat = `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var idFlat int64
	if err := tx.QueryRowContext(ctx, InsertCreateFlat, newFlat.BuildingID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment).Scan(&idFlat); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateFlatMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateFlatMethod)
	return idFlat, nil
}

// SelectImages select list images for advert
func (r *AdvertRepo) SelectImages(ctx context.Context, advertId int64) ([]*models.ImageResp, error) {
	// QuerySelectImages = `SELECT id, photo, priority FROM image WHERE advert_id = $1 AND is_deleted = false`
	rows, err := r.db.Query(QuerySelectImages, advertId)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
		return nil, err
	}
	defer rows.Close()

	images := []*models.ImageResp{}

	for rows.Next() {
		var id int64
		var photo string
		var priority int
		if err := rows.Scan(&id, &photo, &priority); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
			return nil, err
		}
		image := &models.ImageResp{
			ID:       id,
			Photo:    photo,
			Priority: priority,
		}
		images = append(images, image)
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod)
	return images, nil
}

// SelectPriceChanges select list priceChanges for advert
func (r *AdvertRepo) SelectPriceChanges(ctx context.Context, advertId int64) ([]*models.PriceChangeData, error) {
	// QuerySelectPriceChanges = `SELECT price, created_at FROM price_change WHERE advert_id = $1 AND is_deleted = false`
	rows, err := r.db.Query(QuerySelectPriceChanges, advertId)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
		return nil, err
	}
	defer rows.Close()

	prices := []*models.PriceChangeData{}

	for rows.Next() {
		var price int64
		var data time.Time
		if err := rows.Scan(&price, &data); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
			return nil, err
		}
		priceChange := &models.PriceChangeData{
			Price:        price,
			DateCreation: data,
		}
		prices = append(prices, priceChange)
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod)
	return prices, nil
}

// GetTypeAdvertById return type of advert
func (r *AdvertRepo) GetTypeAdvertById(ctx context.Context, id int64) (*models.AdvertTypeAdvert, error) {
	/*
			QueryGetTypeAdvertById = `SELECT                   CASE
			WHEN ath.house_id IS NOT NULL THEN 'House'
			WHEN atf.flat_id IS NOT NULL THEN 'Flat'
			ELSE 'None'
		END AS type_advert FROM advert AS a LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id WHERE a.id=$1`
	*/
	res := r.db.QueryRowContext(ctx, QueryGetTypeAdvertById, id)

	var advertType *models.AdvertTypeAdvert

	if err := res.Scan(&advertType); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetTypeAdvertByIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetTypeAdvertByIdMethod)
	return advertType, nil
}

// GetHouseAdvertById retrieves full information about house advert from the database.
func (r *AdvertRepo) GetHouseAdvertById(ctx context.Context, id int64) (*models.AdvertData, error) {
	/*
			QueryGetHouseAdvertById = `
			SELECT
		        a.id,
		        a.type_placement,
		        a.title,
		        a.description,
		        pc.price,
		        a.phone,
		        a.is_agent,
				ad.metro,
				hn.name,
				s.name,
				t.name,
				p.name,
				ST_AsText(ad.address_point::geometry),
		        h.ceiling_height,
		        h.square_area,
		        h.square_house,
		        h.bedroom_count,
		        h.status_area_house,
		        h.cottage,
		        h.status_home_house,
		        b.floor,
		        b.year_creation,
		        COALESCE(b.material_building, 'Brick') as material,
		        a.created_at,
				CASE
					WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
					ELSE false
				END AS is_liked,
				CASE
					WHEN sva.advert_id IS NOT NULL THEN true
					ELSE false
				END AS is_viewed,
		        cx.id AS complexid,
		        c.photo AS companyphoto,
		        c.name AS companyname,
		        cx.name AS complexname
		    FROM
		        advert AS a
		    JOIN
		        advert_type_house AS at ON a.id = at.advert_id
		    JOIN
		        house AS h ON h.id = at.house_id
		    JOIN
		        building AS b ON h.building_id = b.id
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
			LEFT JOIN
				favourite_advert AS fa ON fa.advert_id=a.id AND fa.user_id=$2
			LEFT JOIN
				statistic_view_advert AS sva ON sva.advert_id=a.id AND sva.user_id=$2
		    LEFT JOIN
		        complex AS cx ON b.complex_id = cx.id
		    LEFT JOIN
		        company AS c ON cx.company_id = c.id
		    JOIN
		        LATERAL (
		            SELECT *
		            FROM price_change AS pc
		            WHERE pc.advert_id = a.id
		            ORDER BY pc.created_at DESC
		            LIMIT 1
		        ) AS pc ON TRUE
		    WHERE
		        a.id = $1 AND a.is_deleted = FALSE;`

	*/

	userId, ok := ctx.Value(middleware.CookieName).(int64)

	if !ok {
		userId = 0
	}

	res := r.db.QueryRowContext(ctx, QueryGetHouseAdvertById, id, userId)

	advertData := &models.AdvertData{}
	var cottage, isViewed bool
	var squareHouse, squareArea, ceilingHeight float64
	var bedroomCount, floor int
	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse
	var complexId, companyPhoto, companyName, complexName sql.NullString
	var metro, houseName, street, town, province string

	if err := res.Scan(
		&advertData.ID,
		&advertData.TypeSale,
		&advertData.Title,
		&advertData.Description,
		&advertData.Price,
		&advertData.Phone,
		&advertData.IsAgent,
		&metro,
		&houseName,
		&street,
		&town,
		&province,
		&advertData.AddressPoint,
		&ceilingHeight,
		&squareArea,
		&squareHouse,
		&bedroomCount,
		&statusArea,
		&cottage,
		&statusHome,
		&floor,
		&advertData.YearCreation,
		&advertData.Material,
		&advertData.DateCreation,
		&advertData.IsLiked,
		&isViewed,
		&complexId,
		&companyPhoto,
		&companyName,
		&complexName,
	); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetHouseAdvertByIdMethod, err)
		return nil, err
	}

	if !isViewed && userId != 0 {
		if err := r.CreateView(ctx, id, userId); err != nil {
			return nil, err
		}
	}

	advertData.AdvertType = "House"

	advertData.HouseProperties = &models.HouseProperties{}
	advertData.HouseProperties.CeilingHeight = ceilingHeight
	advertData.HouseProperties.SquareArea = squareArea
	advertData.HouseProperties.SquareHouse = squareHouse
	advertData.HouseProperties.BedroomCount = bedroomCount
	advertData.HouseProperties.StatusArea = statusArea
	advertData.HouseProperties.Cottage = cottage
	advertData.HouseProperties.StatusHome = statusHome
	advertData.HouseProperties.Floor = floor

	advertData.Address = province + ", " + town + ", " + street + ", " + houseName
	advertData.Metro = metro

	if complexId.String != "" {
		advertData.ComplexProperties = &models.ComplexAdvertProperties{}
		advertData.ComplexProperties.ComplexId = complexId.String
		advertData.ComplexProperties.PhotoCompany = companyPhoto.String
		advertData.ComplexProperties.NameCompany = companyName.String
		advertData.ComplexProperties.NameComplex = complexName.String
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetHouseAdvertByIdMethod)
	return advertData, nil
}

// CheckExistsFlat check exists flat.
func (r *AdvertRepo) CheckExistsFlat(ctx context.Context, advertId int64) (*models.Flat, error) {
	// QueryCheckExistsFlat = `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`

	flat := &models.Flat{}

	res := r.db.QueryRowContext(ctx, QueryCheckExistsFlat, advertId)

	if err := res.Scan(&flat.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsFlatMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsFlatMethod)
	return flat, nil
}

// CheckExistsHouse check exists flat.
func (r *AdvertRepo) CheckExistsHouse(ctx context.Context, advertId int64) (*models.House, error) {
	// QueryCheckExistsHouse = `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`

	house := &models.House{}

	res := r.db.QueryRowContext(ctx, QueryCheckExistsHouse, advertId)

	if err := res.Scan(&house.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsHouseMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsHouseMethod)
	return house, nil
}

// DeleteFlatAdvertById deletes a flat advert by ID.
func (r *AdvertRepo) DeleteFlatAdvertById(ctx context.Context, tx models.Transaction, advertId int64) error {
	/*
			queryGetIdTablesDeleteFlatAdvert = `
		        SELECT
		            f.id as flatid
		        FROM
		            advert AS a
		        JOIN
		            advert_type_flat AS at ON a.id = at.advert_id
		        JOIN
		            flat AS f ON f.id = at.flat_id
		        WHERE a.id=$1;`
	*/
	res := tx.QueryRowContext(ctx, QueryGetIdTablesDeleteFlatAdvert, advertId)

	var flatId int64
	if err := res.Scan(&flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}

	// queryDeleteAdvertByIdDeleteFlatAdvert = `UPDATE advert SET is_deleted=true WHERE id=$1;`
	// queryDeleteAdvertTypeByIdDeleteFlatAdvert = `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`
	// queryDeleteFlatByIdDeleteFlatAdvert = `UPDATE flat SET is_deleted=true WHERE id=$1;`
	// queryDeletePriceChangesDeleteFlatAdvert = `UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`
	// queryDeleteImagesDeleteFlatAdvert = `UPDATE image SET is_deleted=true WHERE advert_id=$1;`

	if _, err := tx.Exec(QueryDeleteAdvertByIdDeleteFlatAdvert, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeleteAdvertTypeByIdDeleteFlatAdvert, advertId, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeleteFlatByIdDeleteFlatAdvert, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeletePriceChangesDeleteFlatAdvert, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeleteImagesDeleteFlatAdvert, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod)
	return nil
}

// DeleteHouseAdvertById deletes a house advert by ID.
func (r *AdvertRepo) DeleteHouseAdvertById(ctx context.Context, tx models.Transaction, advertId int64) error {
	/*
			QueryGetIdTablesDeleteHouseAdvert = `
		        SELECT
		            h.id as houseid
		        FROM
		            advert AS a
		        JOIN
		            advert_type_house AS at ON a.id = at.advert_id
		        JOIN
		            house AS h ON h.id = at.house_id
		        WHERE a.id=$1;`

	*/

	res := tx.QueryRowContext(ctx, QueryGetIdTablesDeleteHouseAdvert, advertId)

	var houseId int64
	if err := res.Scan(&houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}

	// QueryDeleteAdvertByIdDeleteHouseAdvert = `UPDATE advert SET is_deleted=true WHERE id=$1;`
	// QueryDeleteAdvertTypeByIdDeleteHouseAdvert = `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
	// QueryDeleteHouseByIdDeleteHouseAdvert = `UPDATE house SET is_deleted=true WHERE id=$1;`
	// QueryDeletePriceChangesDeleteHouseAdvert = `UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`
	// QueryDeleteImagesDeleteHouseAdvert = `UPDATE image SET is_deleted=true WHERE advert_id=$1;`

	if _, err := tx.Exec(QueryDeleteAdvertByIdDeleteHouseAdvert, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeleteAdvertTypeByIdDeleteHouseAdvert, advertId, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeleteHouseByIdDeleteHouseAdvert, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeletePriceChangesDeleteHouseAdvert, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryDeleteImagesDeleteHouseAdvert, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod)
	return nil
}

// ChangeTypeAdvert Когда мы захотели поменять тип объявления(Дом, Квартира), Меняем сущности в бд
func (r *AdvertRepo) ChangeTypeAdvert(ctx context.Context, tx models.Transaction, advertId int64) (err error) {
	/*
					QueryChangeTypeAdvert = `SELECT 			CASE
					WHEN ath.house_id IS NOT NULL THEN 'House'
					WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					ELSE 'None'
				END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`

			QuerySelectBuildingIdByFlatChangeTypeAdvert = `SELECT b.id AS buildingid, f.id AS flatid  FROM advert AS a JOIN advert_type_flat AS at ON at.advert_id=a.id JOIN flat AS f ON f.id=at.flat_id JOIN building AS b ON f.building_id=b.id WHERE a.id=$1`

			QuerySelectBuildingIdByHouseChangeTypeAdvert = `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`

		QueryInsertFlatChangeTypeAdvert = `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament)
					VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

		QueryInsertHouseChangeTypeAdvert = `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

		QueryInsertTypeFlatChangeTypeAdvert = `INSERT INTO advert_type_flat (advert_id, flat_id) VALUES ($1, $2);`

		QueryInsertTypeHouseChangeTypeAdvert = `INSERT INTO advert_type_house (advert_id, house_id) VALUES ($1, $2);`

		QueryRestoreFlatByIdChangeTypeAdvert = `UPDATE flat SET is_deleted=false WHERE id=$1;`

		QueryRestoreHouseByIdChangeTypeAdvert = `UPDATE house SET is_deleted=false WHERE id=$1;`

		QueryDeleteFlatByIdChangeTypeAdvert = `UPDATE flat SET is_deleted=true WHERE id=$1;`

		QueryDeleteHouseByIdChangeTypeAdvert = `UPDATE house SET is_deleted=true WHERE id=$1;`

		QueryDeleteAdvertTypeFlatChangeTypeAdvert = `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`

		QueryDeleteAdvertTypeHouseChangeTypeAdvert = `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`

		QueryRestoreAdvertTypeFlatChangeTypeAdvert = `UPDATE advert_type_flat SET is_deleted=false WHERE advert_id=$1 AND flat_id=$2;`

		QueryRestoreAdvertTypeHouseChangeTypeAdvert = `UPDATE advert_type_house SET is_deleted=false WHERE advert_id=$1 AND house_id=$2;`
	*/

	var advertType models.AdvertTypeAdvert
	res := r.db.QueryRowContext(ctx, QueryChangeTypeAdvert, advertId)

	if err := res.Scan(&advertType); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
		return err
	}
	var buildingId int64
	switch advertType {
	case models.AdvertTypeFlat:
		res := r.db.QueryRowContext(ctx, QuerySelectBuildingIdByFlatChangeTypeAdvert, advertId)

		var flatId int64

		if err := res.Scan(&buildingId, &flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(QueryDeleteFlatByIdChangeTypeAdvert, flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(QueryDeleteAdvertTypeFlatChangeTypeAdvert, advertId, flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		house, err := r.CheckExistsHouse(ctx, advertId)
		if err != nil {
			var id int64
			house := &models.House{}
			err := tx.QueryRowContext(ctx, QueryInsertHouseChangeTypeAdvert, buildingId, house.CeilingHeight, house.SquareArea, house.SquareHouse, house.BedroomCount, models.StatusAreaDNP, house.Cottage, models.StatusHomeCompleteNeed).Scan(&id)
			if err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
			if _, err := tx.Exec(QueryInsertTypeHouseChangeTypeAdvert, advertId, id); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		} else {
			if _, err := tx.Exec(QueryRestoreHouseByIdChangeTypeAdvert, house.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}

			if _, err := tx.Exec(QueryRestoreAdvertTypeHouseChangeTypeAdvert, advertId, house.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		}
	case models.AdvertTypeHouse:
		res := r.db.QueryRowContext(ctx, QuerySelectBuildingIdByHouseChangeTypeAdvert, advertId)

		var houseId int64

		if err := res.Scan(&buildingId, &houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(QueryDeleteHouseByIdChangeTypeAdvert, houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(QueryDeleteAdvertTypeHouseChangeTypeAdvert, advertId, houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		flat, err := r.CheckExistsFlat(ctx, advertId)
		if err != nil {
			var id int64
			flat = &models.Flat{}
			err := tx.QueryRowContext(ctx, QueryInsertFlatChangeTypeAdvert, buildingId, flat.Floor, flat.CeilingHeight, flat.SquareGeneral, flat.RoomCount, flat.SquareResidential, flat.Apartment).Scan(&id)
			if err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
			if _, err := tx.Exec(QueryInsertTypeFlatChangeTypeAdvert, advertId, id); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		} else {
			if _, err := tx.Exec(QueryRestoreFlatByIdChangeTypeAdvert, flat.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}

			if _, err := tx.Exec(QueryRestoreAdvertTypeFlatChangeTypeAdvert, advertId, flat.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod)
	return nil
}

// UpdateHouseAdvertById updates a house advert in the database.
func (r *AdvertRepo) UpdateHouseAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error {
	/*
			QueryGetIdTablesUpdateHouseAdvert = `
		        SELECT
		            b.id as buildingid,
		            h.id as houseid,
		            pc.price
		        FROM
		            advert AS a
		        JOIN
		            advert_type_house AS at ON a.id = at.advert_id
		        JOIN
		            house AS h ON h.id = at.house_id
		        JOIN
		            building AS b ON h.building_id = b.id
		        LEFT JOIN
		            LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		        WHERE a.id=$1;`
	*/

	res := tx.QueryRowContext(ctx, QueryGetIdTablesUpdateHouseAdvert, advertUpdateData.ID)

	var buildingId, houseId int64
	var price float64
	if err := res.Scan(&buildingId, &houseId, &price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}

	id, err := r.CreateProvince(ctx, tx, advertUpdateData.Address.Province)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateTown(ctx, tx, id, advertUpdateData.Address.Town)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateStreet(ctx, tx, id, advertUpdateData.Address.Street)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateHouseAddress(ctx, tx, id, advertUpdateData.Address.House)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateAddress(ctx, tx, id, advertUpdateData.Address.Metro, advertUpdateData.Address.AddressPoint)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	// QueryUpdateAdvertByIdUpdateHouseAdvert = `UPDATE advert SET type_placement=$1, title=$2, description=$3, phone=$4, is_agent=$5 WHERE id=$6;`
	// QueryUpdateBuildingByIdUpdateHouseAdvert = `UPDATE building SET floor=$1, material_building=$2, address_id=$3, year_creation=$4 WHERE id=$5;`
	// QueryUpdateHouseByIdUpdateHouseAdvert = `UPDATE house SET ceiling_height=$1, square_area=$2, square_house=$3, bedroom_count=$4, status_area_house=$5, cottage=$6, status_home_house=$7 WHERE id=$8;`

	if _, err := tx.Exec(QueryUpdateAdvertByIdUpdateHouseAdvert, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryUpdateBuildingByIdUpdateHouseAdvert, advertUpdateData.HouseProperties.Floor, advertUpdateData.Material, id, advertUpdateData.YearCreation, buildingId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryUpdateHouseByIdUpdateHouseAdvert, advertUpdateData.HouseProperties.CeilingHeight, advertUpdateData.HouseProperties.SquareArea, advertUpdateData.HouseProperties.SquareHouse, advertUpdateData.HouseProperties.BedroomCount, advertUpdateData.HouseProperties.StatusArea, advertUpdateData.HouseProperties.Cottage, advertUpdateData.HouseProperties.StatusHome, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if advertUpdateData.Price != price {
		/*
					QueryInsertPriceChangeUpdateHouseAdvert = `INSERT INTO price_change (advert_id, price)
			            VALUES ($1, $2)`
		*/
		if _, err := tx.Exec(QueryInsertPriceChangeUpdateHouseAdvert, advertUpdateData.ID, advertUpdateData.Price); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
			return err
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod)
	return nil
}

// UpdateFlatAdvertById updates a flat advert in the database.
func (r *AdvertRepo) UpdateFlatAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error {
	/*
			QueryGetIdTablesUpdateFlatAdvert = `
		        SELECT
		            b.id as buildingid,
		            f.id as flatid,
		            pc.price
		        FROM
		            advert AS a
		        JOIN
		            advert_type_flat AS at ON a.id = at.advert_id
		        JOIN
		            flat AS f ON f.id = at.flat_id
		        JOIN
		            building AS b ON f.building_id = b.id
		        LEFT JOIN
		            LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		        WHERE a.id=$1;`
	*/

	res := tx.QueryRowContext(ctx, QueryGetIdTablesUpdateFlatAdvert, advertUpdateData.ID)

	var buildingId, flatId int64
	var price float64
	if err := res.Scan(&buildingId, &flatId, &price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}

	id, err := r.CreateProvince(ctx, tx, advertUpdateData.Address.Province)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateTown(ctx, tx, id, advertUpdateData.Address.Town)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateStreet(ctx, tx, id, advertUpdateData.Address.Street)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateHouseAddress(ctx, tx, id, advertUpdateData.Address.House)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateAddress(ctx, tx, id, advertUpdateData.Address.Metro, advertUpdateData.Address.AddressPoint)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	// QueryUpdateAdvertByIdUpdateFlatAdvert = `UPDATE advert SET type_placement=$1, title=$2, description=$3, phone=$4, is_agent=$5 WHERE id=$6;`
	// QueryUpdateBuildingByIdUpdateFlatAdvert = `UPDATE building SET floor=$1, material_building=$2, address_id=$3, year_creation=$4 WHERE id=$5;`
	// QueryUpdateFlatByIdUpdateFlatAdvert = `UPDATE flat SET floor=$1, ceiling_height=$2, square_general=$3, bedroom_count=$4, square_residential=$5, apartament=$6 WHERE id=$7;`

	if _, err := tx.Exec(QueryUpdateAdvertByIdUpdateFlatAdvert, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryUpdateBuildingByIdUpdateFlatAdvert, advertUpdateData.FlatProperties.FloorGeneral, advertUpdateData.Material, id, advertUpdateData.YearCreation, buildingId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(QueryUpdateFlatByIdUpdateFlatAdvert, advertUpdateData.FlatProperties.Floor, advertUpdateData.FlatProperties.CeilingHeight, advertUpdateData.FlatProperties.SquareGeneral, advertUpdateData.FlatProperties.RoomCount, advertUpdateData.FlatProperties.SquareResidential, advertUpdateData.FlatProperties.Apartment, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}

	if advertUpdateData.Price != price {
		/*
					QueryInsertPriceChangeUpdateFlatAdvert = `INSERT INTO price_change (advert_id, price)
			            VALUES ($1, $2)`

		*/
		if _, err := tx.Exec(QueryInsertPriceChangeUpdateFlatAdvert, advertUpdateData.ID, advertUpdateData.Price); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
			return err
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod)
	return nil
}

// GetFlatAdvertById retrieves full information about flat advert from the database.
func (r *AdvertRepo) GetFlatAdvertById(ctx context.Context, id int64) (*models.AdvertData, error) {
	/*
			QueryGetFlatAdvertById = `
			SELECT
				a.id,
				a.type_placement,
				a.title,
				a.description,
				pc.price,
				a.phone,
				a.is_agent,
				ad.metro,
				hn.name,
				s.name,
				t.name,
				p.name,
				ST_AsText(ad.address_point::geometry),
		        f.floor,
		        f.ceiling_height,
		        f.square_general,
		        f.bedroom_count,
		        f.square_residential,
		        f.apartament,
		        b.floor AS floorGeneral,
		        b.year_creation,
		        COALESCE(b.material_building, 'Brick') as material,
		        a.created_at,
				CASE
					WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
					ELSE false
				END AS is_liked,
				CASE
					WHEN sva.advert_id IS NOT NULL THEN true
					ELSE false
				END AS is_viewed,
		        cx.id AS complexid,
		        c.photo AS companyphoto,
		        c.name AS companyname,
		        cx.name AS complexname
		    FROM
		        advert AS a
		    JOIN
		        advert_type_flat AS at ON a.id = at.advert_id
		    JOIN
		        flat AS f ON f.id = at.flat_id
		    JOIN
		        building AS b ON f.building_id = b.id
				JOIN address AS ad ON b.address_id=ad.id
				JOIN house_name AS hn ON hn.id=ad.house_name_id
				JOIN street AS s ON s.id=hn.street_id
				JOIN town AS t ON t.id=s.town_id
				JOIN province AS p ON p.id=t.province_id
			LEFT JOIN
				favourite_advert AS fa ON fa.advert_id=a.id AND fa.user_id=$2
			LEFT JOIN
				statistic_view_advert AS sva ON sva.advert_id=a.id AND sva.user_id=$2
		    LEFT JOIN
		        complex AS cx ON b.complex_id = cx.id
		    LEFT JOIN
		        company AS c ON cx.company_id = c.id
		    LEFT JOIN
		        LATERAL (
		            SELECT *
		            FROM price_change AS pc
		            WHERE pc.advert_id = a.id
		            ORDER BY pc.created_at DESC
		            LIMIT 1
		        ) AS pc ON TRUE
		    WHERE
		        a.id = $1 AND a.is_deleted = FALSE;`

	*/

	userId, ok := ctx.Value(middleware.CookieName).(int64)

	if !ok {
		userId = 0
	}

	res := r.db.QueryRowContext(ctx, QueryGetFlatAdvertById, id, userId)

	advertData := &models.AdvertData{}
	var floor, floorGeneral, roomCount int
	var squareGenereal, squareResidential, ceilingHeight float64
	var apartament sql.NullBool
	var isViewed bool
	var complexId, companyPhoto, companyName, complexName sql.NullString
	var metro, houseName, street, town, province string

	if err := res.Scan(
		&advertData.ID,
		&advertData.TypeSale,
		&advertData.Title,
		&advertData.Description,
		&advertData.Price,
		&advertData.Phone,
		&advertData.IsAgent,
		&metro,
		&houseName,
		&street,
		&town,
		&province,
		&advertData.AddressPoint,
		&floor,
		&ceilingHeight,
		&squareGenereal,
		&roomCount,
		&squareResidential,
		&apartament,
		&floorGeneral,
		&advertData.YearCreation,
		&advertData.Material,
		&advertData.DateCreation,
		&advertData.IsLiked,
		&isViewed,
		&complexId,
		&companyPhoto,
		&companyName,
		&complexName,
	); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetFlatAdvertByIdMethod, err)
		return nil, err
	}

	if !isViewed && userId != 0 {
		if err := r.CreateView(ctx, id, userId); err != nil {
			return nil, err
		}
	}

	advertData.AdvertType = "Flat"
	advertData.FlatProperties = &models.FlatProperties{}
	advertData.FlatProperties.CeilingHeight = ceilingHeight
	advertData.FlatProperties.Apartment = apartament.Bool
	advertData.FlatProperties.SquareResidential = squareResidential
	advertData.FlatProperties.RoomCount = roomCount
	advertData.FlatProperties.SquareGeneral = squareGenereal
	advertData.FlatProperties.FloorGeneral = floorGeneral
	advertData.FlatProperties.Floor = floor

	advertData.Address = province + ", " + town + ", " + street + ", " + houseName
	advertData.Metro = metro

	if complexId.String != "" {
		advertData.ComplexProperties = &models.ComplexAdvertProperties{}
		advertData.ComplexProperties.ComplexId = complexId.String
		advertData.ComplexProperties.PhotoCompany = companyPhoto.String
		advertData.ComplexProperties.NameCompany = companyName.String
		advertData.ComplexProperties.NameComplex = complexName.String
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetFlatAdvertByIdMethod)
	return advertData, nil
}

// GetSquareAdverts retrieves square adverts from the database.
func (r *AdvertRepo) GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error) {
	/*
			QueryBaseAdvertGetSquareAdverts = `
		        SELECT
		            a.id,
		            a.type_placement,
					CASE
		           		WHEN ath.house_id IS NOT NULL THEN 'House'
		           		WHEN atf.flat_id IS NOT NULL THEN 'Flat'
		           		ELSE 'None'
		       		END AS type_advert,
		            i.photo,
		            pc.price,
		            a.created_at
		        FROM
		            advert AS a
					LEFT JOIN advert_type_house AS ath ON ath.advert_id=a.id
					LEFT JOIN advert_type_flat AS atf ON atf.advert_id=a.id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		        ORDER BY
		            a.created_at DESC
		        LIMIT $1
		        OFFSET $2;`

			QueryFlatGetSquareAdverts = `
		        SELECT
		            f.square_general,
		            f.floor,
		            ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral,
		            f.bedroom_count
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id=at.flat_id
		            JOIN building AS b ON f.building_id=b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id=$1 AND a.is_deleted = FALSE
		        ORDER BY
		            a.created_at DESC;`
			QueryHouseGetSquareAdverts = `
		        	SELECT
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            h.bedroom_count,
		            b.floor
		        FROM
		            advert AS a
		            JOIN advert_type_house AS at ON a.id = at.advert_id
		            JOIN house AS h ON h.id=at.house_id
		            JOIN building AS b ON h.building_id=b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id=$1
		        ORDER BY
		            a.created_at DESC;`

	*/

	rows, err := r.db.Query(QueryBaseAdvertGetSquareAdverts, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeSale, &squareAdvert.TypeAdvert, &squareAdvert.Photo, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
			return nil, err
		}
		var metro, province, town, street, houseName string
		switch squareAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral, roomCount int
			row := r.db.QueryRowContext(ctx, QueryFlatGetSquareAdverts, squareAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral, &roomCount); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
				return nil, err
			}
			squareAdvert.FlatProperties = &models.FlatSquareProperties{}
			squareAdvert.FlatProperties.Floor = floor
			squareAdvert.FlatProperties.FloorGeneral = floorGeneral
			squareAdvert.FlatProperties.RoomCount = roomCount
			squareAdvert.FlatProperties.SquareGeneral = squareGeneral
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var bedroomCount, floor int
			row := r.db.QueryRowContext(ctx, QueryHouseGetSquareAdverts, squareAdvert.ID)
			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &bedroomCount, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
				return nil, err
			}
			squareAdvert.HouseProperties = &models.HouseSquareProperties{}
			squareAdvert.HouseProperties.Cottage = cottage
			squareAdvert.HouseProperties.SquareHouse = squareHouse
			squareAdvert.HouseProperties.SquareArea = squareArea
			squareAdvert.HouseProperties.BedroomCount = bedroomCount
			squareAdvert.HouseProperties.Floor = floor
		}

		squareAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		squareAdvert.Metro = metro

		squareAdverts = append(squareAdverts, squareAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod)
	return squareAdverts, nil
}

// GetRectangleAdverts retrieves rectangle adverts from the database with search.
func (r *AdvertRepo) GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error) {
	/*
			QueryBaseAdvertGetRectangleAdverts = `
		        SELECT
					a.id,
					a.title,
					a.description,
					CASE
					   WHEN ath.house_id IS NOT NULL THEN 'House'
					   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					   ELSE 'None'
				    END AS type_advert,
		            CASE
		                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
		                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
		                ELSE 0
		            END AS rcount,
		            a.phone,
		            a.type_placement,
		            pc.price,
		            i.photo,
		            a.created_at
		        FROM
		            advert AS a
		            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
					LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
		            LEFT JOIN flat AS f ON f.id = atf.flat_id
		            LEFT JOIN house AS h ON h.id = ath.house_id
		            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		            AND a.is_deleted = FALSE
		            AND pc.price >= $1
		            AND pc.price <= $2
		            AND CONCAT_WS(', ', COALESCE(p.name, ''), COALESCE(t.name, ''), COALESCE(s.name, ''), COALESCE(hn.name, '')) ILIKE $3`
			QueryFlatGetRectangleAdverts = `
		        SELECT
		            f.square_general,
		            f.floor,
					ST_AsText(ad.address_point::geometry),
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id = at.flat_id
		            JOIN building AS b ON f.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
			QueryHouseGetRectangleAdverts = `
		        SELECT
					ST_AsText(ad.address_point::geometry),
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            b.floor
		        FROM
					advert AS a
					JOIN advert_type_house AS at ON a.id = at.advert_id
					JOIN house AS h ON h.id = at.house_id
					JOIN building AS b ON h.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
	*/
	QueryBaseAdvertGetRectangleAdverts := `
		        SELECT
					a.id,
					a.title,
					a.description,
					CASE
					   WHEN ath.house_id IS NOT NULL THEN 'House'
					   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					   ELSE 'None'
				    END AS type_advert,
		            CASE
		                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
		                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
		                ELSE 0
		            END AS rcount,
		            a.phone,
		            a.type_placement,
		            pc.price,
		            i.photo,
		            a.created_at
		        FROM
		            advert AS a
		            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
					LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
		            LEFT JOIN flat AS f ON f.id = atf.flat_id
		            LEFT JOIN house AS h ON h.id = ath.house_id
		            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		            AND a.is_deleted = FALSE
		            AND pc.price >= $1
		            AND pc.price <= $2
		            AND CONCAT_WS(', ', COALESCE(p.name, ''), COALESCE(t.name, ''), COALESCE(s.name, ''), COALESCE(hn.name, '')) ILIKE $3`

	pageInfo := &models.PageInfo{}

	var argsForQuery []interface{}
	i := 4 // Изначально в запросе проставлены minPrice, maxPrice и address, поэтому начинаем с 4 для формирования поиска

	advertFilter.Address = "%" + advertFilter.Address + "%"

	if advertFilter.DealType != "" {
		QueryBaseAdvertGetRectangleAdverts += " AND a.type_placement = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.DealType)
		i++
	}

	if advertFilter.AdvertType != "" {
		QueryBaseAdvertGetRectangleAdverts = "SELECT * FROM (" + QueryBaseAdvertGetRectangleAdverts + ") AS subqueryforadverttypecalculate WHERE type_advert = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.AdvertType)
		i++
	}

	if advertFilter.RoomCount != 0 {
		QueryBaseAdvertGetRectangleAdverts = "SELECT * FROM (" + QueryBaseAdvertGetRectangleAdverts + ") AS subqueryforroomcountcalculate WHERE rcount = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.RoomCount)
		i++
	}

	queryCount := "SELECT COUNT(*) FROM (" + QueryBaseAdvertGetRectangleAdverts + ") AS subqueryforpaginate"
	QueryBaseAdvertGetRectangleAdverts += " ORDER BY created_at DESC LIMIT $" + fmt.Sprint(i) + " OFFSET $" + fmt.Sprint(i+1) + ";"
	rowCountQuery := r.db.QueryRowContext(ctx, queryCount, append([]interface{}{advertFilter.MinPrice, advertFilter.MaxPrice, advertFilter.Address}, argsForQuery...)...)

	if err := rowCountQuery.Scan(&pageInfo.TotalElements); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
		return nil, err
	}

	argsForQuery = append(argsForQuery, advertFilter.Page, advertFilter.Offset)
	rows, err := r.db.Query(QueryBaseAdvertGetRectangleAdverts, append([]interface{}{advertFilter.MinPrice, advertFilter.MaxPrice, advertFilter.Address}, argsForQuery...)...)

	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
		return nil, err
	}

	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
			return nil, err
		}

		var metro, houseName, street, town, province string

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, QueryBaseAdvertGetRectangleAdverts, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.AddressPoint, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, QueryHouseGetRectangleAdverts, rectangleAdvert.ID)

			if err := row.Scan(&rectangleAdvert.AddressPoint, &metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}

	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
		return nil, err
	}

	pageInfo.PageSize = advertFilter.Page
	pageInfo.TotalPages = pageInfo.TotalElements / pageInfo.PageSize

	if pageInfo.TotalElements%pageInfo.PageSize != 0 {
		pageInfo.TotalPages++
	}

	pageInfo.CurrentPage = (advertFilter.Offset / pageInfo.PageSize) + 1

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod)

	return &models.AdvertDataPage{
		Adverts:  rectangleAdverts,
		PageInfo: pageInfo,
	}, nil
}

// GetRectangleAdvertsByUserId retrieves rectangle adverts from the database by user id.
func (r *AdvertRepo) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId int64) ([]*models.AdvertRectangleData, error) {
	/*
			QueryBaseAdvertGetRectangleAdvertsByUser = `
		        SELECT
					a.id,
					a.title,
					a.description,
					CASE
					   WHEN ath.house_id IS NOT NULL THEN 'House'
					   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					   ELSE 'None'
				    END AS type_advert,
		            CASE
		                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
		                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
		                ELSE 0
		            END AS rcount,
		            a.phone,
		            a.type_placement,
		            pc.price,
		            i.photo,
		            a.created_at,
					CASE
						WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
						ELSE false
					END AS is_liked
		        FROM
		            advert AS a
		            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
					LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
					LEFT JOIN favourite_advert AS fa ON a.id=fa.advert_id AND fa.user_id=$1
		            LEFT JOIN flat AS f ON f.id = atf.flat_id
		            LEFT JOIN house AS h ON h.id = ath.house_id
		            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		            AND a.is_deleted = FALSE
		            AND a.user_id = $1
					ORDER BY a.created_at DESC
					LIMIT $2
					OFFSET $3`
			QueryFlatGetRectangleAdvertsByUser = `
		        SELECT
		            f.square_general,
		            f.floor,
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id = at.flat_id
		            JOIN building AS b ON f.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
			QueryHouseGetRectangleAdvertsByUser = `
		        SELECT
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            b.floor
		        FROM
					advert AS a
					JOIN advert_type_house AS at ON a.id = at.advert_id
					JOIN house AS h ON h.id = at.house_id
					JOIN building AS b ON h.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
	*/

	rows, err := r.db.Query(QueryBaseAdvertGetRectangleAdvertsByUser, userId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var metro, houseName, street, town, province string
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert,
			&roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price,
			&rectangleAdvert.Photo, &rectangleAdvert.DateCreation, &rectangleAdvert.IsLiked)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, QueryFlatGetRectangleAdvertsByUser, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, QueryHouseGetRectangleAdvertsByUser, rectangleAdvert.ID)

			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return rectangleAdverts, nil
}

// GetRectangleAdvertsByComplexId retrieves rectangle adverts from the database by complex id.
func (r *AdvertRepo) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId int64) ([]*models.AdvertRectangleData, error) {
	/*
			QueryBaseAdvertGetRectangleAdvertsByComplex = `
		        SELECT
					a.id,
					a.title,
					a.description,
					CASE
					   WHEN ath.house_id IS NOT NULL THEN 'House'
					   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					   ELSE 'None'
				    END AS type_advert,
		            CASE
		                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
		                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
		                ELSE 0
		            END AS rcount,
		            a.phone,
		            a.type_placement,
		            pc.price,
		            i.photo,
		            a.created_at
		        FROM
		            advert AS a
		            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
					LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
		            LEFT JOIN flat AS f ON f.id = atf.flat_id
		            LEFT JOIN house AS h ON h.id = ath.house_id
		            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		            AND a.is_deleted = FALSE
		            AND b.complex_id = $1
					ORDER BY a.created_at DESC
					LIMIT $2
					OFFSET $3`
			QueryFlatGetRectangleAdvertsByComplex := `
		        SELECT
		            f.square_general,
		            f.floor,
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id = at.flat_id
		            JOIN building AS b ON f.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
			QueryHouseGetRectangleAdvertsByComplex := `
		        SELECT
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            b.floor
		        FROM
					advert AS a
					JOIN advert_type_house AS at ON a.id = at.advert_id
					JOIN house AS h ON h.id = at.house_id
					JOIN building AS b ON h.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`

	*/

	rows, err := r.db.Query(QueryBaseAdvertGetRectangleAdvertsByComplex, complexId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var metro, houseName, street, town, province string
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert,
			&roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price,
			&rectangleAdvert.Photo, &rectangleAdvert.DateCreation)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, QueryFlatGetRectangleAdvertsByComplex, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, QueryHouseGetRectangleAdvertsByComplex, rectangleAdvert.ID)

			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return rectangleAdverts, nil
}

// LikeAdvert creates a like in the database.
func (r *AdvertRepo) LikeAdvert(ctx context.Context, advertId int64, userId int64) error {
	query := `SELECT advert_id, user_id FROM favourite_advert WHERE advert_id = $1 AND user_id = $2`

	res := r.db.QueryRow(query, advertId, userId)

	var adId, usId int64
	if err := res.Scan(&adId, &usId); err == nil {
		update := `UPDATE favourite_advert SET is_deleted = false WHERE advert_id = $1 AND user_id = $2`
		if _, err := r.db.Exec(update, adId, usId); err != nil {
			// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
			return err
		}
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return nil
	}

	insert := `INSERT INTO favourite_advert (advert_id, user_id) VALUES ($1, $2)`
	if _, err := r.db.Exec(insert, advertId, userId); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return nil
}

// DislikeAdvert set dislike in the database.
func (r *AdvertRepo) DislikeAdvert(ctx context.Context, advertId int64, userId int64) error {
	query := `SELECT advert_id, user_id FROM favourite_advert WHERE advert_id = $1 AND user_id = $2`

	res := r.db.QueryRow(query, advertId, userId)

	var adId, usId int64
	if err := res.Scan(&adId, &usId); err == nil {
		update := `UPDATE favourite_advert SET is_deleted = true WHERE advert_id = $1 AND user_id = $2`
		if _, err := r.db.Exec(update, adId, usId); err != nil {
			// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
			return err
		}
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return nil
}

// GetRectangleAdvertsByUserId retrieves rectangle adverts from the database by user id.
func (r *AdvertRepo) GetRectangleAdvertsLikedByUserId(ctx context.Context, pageSize, offset int, userId int64) ([]*models.AdvertRectangleData, error) {
	/*
			QueryBaseAdvertGetRectangleAdvertsLikedByUser = `
		        SELECT
					a.id,
					a.title,
					a.description,
					CASE
					   WHEN ath.house_id IS NOT NULL THEN 'House'
					   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
					   ELSE 'None'
				    END AS type_advert,
		            CASE
		                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
		                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
		                ELSE 0
		            END AS rcount,
		            a.phone,
		            a.type_placement,
		            pc.price,
		            i.photo,
		            a.created_at,
					CASE
						WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
						ELSE false
					END AS is_liked
		        FROM
		            advert AS a
		            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
					LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
					JOIN favourite_advert AS fa ON (a.id=fa.advert_id AND fa.user_id=$1 AND fa.is_deleted = false)
		            LEFT JOIN flat AS f ON f.id = atf.flat_id
		            LEFT JOIN house AS h ON h.id = ath.house_id
		            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		            LEFT JOIN LATERAL (
		                SELECT *
		                FROM price_change AS pc
		                WHERE pc.advert_id = a.id
		                ORDER BY pc.created_at DESC
		                LIMIT 1
		            ) AS pc ON TRUE
		            JOIN image AS i ON i.advert_id = a.id
		        WHERE i.priority = (
		                SELECT MIN(priority)
		                FROM image
		                WHERE advert_id = a.id
		                    AND is_deleted = FALSE
		            )
		            AND i.is_deleted = FALSE
		            AND a.is_deleted = FALSE
		            AND a.user_id = $1
					ORDER BY a.created_at DESC
					LIMIT $2
					OFFSET $3`
			QueryFlatGetRectangleAdvertsLikedByUser := `
		        SELECT
		            f.square_general,
		            f.floor,
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            b.floor AS floorgeneral
		        FROM
		            advert AS a
		            JOIN advert_type_flat AS at ON a.id = at.advert_id
		            JOIN flat AS f ON f.id = at.flat_id
		            JOIN building AS b ON f.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`
			QueryHouseGetRectangleAdvertsLikedByUser := `
		        SELECT
					ad.metro,
					hn.name,
					s.name,
					t.name,
					p.name,
		            h.cottage,
		            h.square_house,
		            h.square_area,
		            b.floor
		        FROM
					advert AS a
					JOIN advert_type_house AS at ON a.id = at.advert_id
					JOIN house AS h ON h.id = at.house_id
					JOIN building AS b ON h.building_id = b.id
					JOIN address AS ad ON b.address_id=ad.id
					JOIN house_name AS hn ON hn.id=ad.house_name_id
					JOIN street AS s ON s.id=hn.street_id
					JOIN town AS t ON t.id=s.town_id
					JOIN province AS p ON p.id=t.province_id
		        WHERE a.id = $1
		        ORDER BY
		            a.created_at DESC;`

	*/

	rows, err := r.db.Query(QueryBaseAdvertGetRectangleAdvertsLikedByUser, userId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var metro, houseName, street, town, province string
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert,
			&roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price,
			&rectangleAdvert.Photo, &rectangleAdvert.DateCreation, &rectangleAdvert.IsLiked)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, QueryFlatGetRectangleAdvertsLikedByUser, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, QueryHouseGetRectangleAdvertsLikedByUser, rectangleAdvert.ID)

			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return rectangleAdverts, nil
}

// SelectCountLikes count likes by advert in the database.
func (r *AdvertRepo) SelectCountLikes(ctx context.Context, id int64) (int64, error) {
	query := `SELECT COUNT (*) FROM favourite_advert WHERE advert_id=$1 AND is_deleted=false;`

	res := r.db.QueryRow(query, id)

	var countLikes int64
	if err := res.Scan(&countLikes); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return countLikes, nil
}

// SelectCountViews count views by advert in the database.
func (r *AdvertRepo) SelectCountViews(ctx context.Context, id int64) (int64, error) {
	query := `SELECT COUNT (*) FROM statistic_view_advert WHERE advert_id=$1`

	res := r.db.QueryRow(query, id)

	var countViews int64
	if err := res.Scan(&countViews); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return countViews, nil
}

// CreateView creates a view in the database.
func (r *AdvertRepo) CreateView(ctx context.Context, advertId int64, userId int64) error {
	query := `SELECT advert_id, user_id FROM statistic_view_advert WHERE advert_id = $1 AND user_id = $2`

	res := r.db.QueryRow(query, advertId, userId)

	var adId, usId int64
	if err := res.Scan(&adId, &usId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return nil
	}

	insert := `INSERT INTO statistic_view_advert (advert_id, user_id) VALUES ($1, $2)`
	if _, err := r.db.Exec(insert, advertId, userId); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return nil
}
