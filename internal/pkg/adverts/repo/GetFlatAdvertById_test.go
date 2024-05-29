package repo_test

/*
import (
	"2024_1_TeaStealers/internal/pkg/adverts/repo"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"regexp"
)

func (suite *AdvertRepoTestSuite) TestGetFlatAdvertById() {
	ctx := context.Background()
	id := int64(1)
	userId := int64(1)

	rows := sqlmock.NewRows([]string{"id", "type_placement", "title", "description", "price", "phone", "is_agent", "metro", "name", "name", "name", "name", "address_point", "floor", "ceiling_height", "square_general", "bedroom_count", "square_residential", "apartament", "floorGeneral", "year_creation", "material", "created_at", "is_liked", "complexid", "companyphoto", "companyname", "complexname"}).
		AddRow(1, "Sale", "Test Title", "Test Description", 100000, "1234567890", true, "Test Metro", "Test House Name", "Test Street", "Test Town", "Test Province", "Test Address Point", 1, 3.0, 100.0, 2, 80.0, true, 5, 2000, "Brick", "2022-01-01T00:00:00Z", true, 1, "test_photo.jpg", "Test Company", "Test Complex")

	query := `
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

	suite.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id, userId).WillReturnRows(rows)
	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(suite.db, logger)
	advertData, err := rep.GetFlatAdvertById(ctx, id)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), advertData)
	// Add more assertions here to check the returned AdvertData
}

*/
