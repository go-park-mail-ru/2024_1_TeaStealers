package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// AdvertRepo represents a repository for adverts changes.
type QuestionRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of AdvertRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *QuestionRepo {
	return &QuestionRepo{db: db, logger: logger}
}

// StoreAnswer creates a new Answer in the database on question.
func (r *QuestionRepo) StoreAnswer(ctx context.Context, newAnswer *models.QuestionAnswer) error {
	insert := `INSERT INTO question_answer (user_id, question_id, mark) VALUES ($1, $2, $3)`
	if _, err := r.db.ExecContext(ctx, insert, newAnswer.UserID, newAnswer.QuestionID, newAnswer.Mark); err != nil {
		//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	//utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// SelectQuestionsByTheme retrieves question from the database by theme.
func (r *QuestionRepo) SelectQuestionsByTheme(ctx context.Context, theme models.QuestionTheme) ([]*models.QuestionResp, error) {
	queryAllAnswersByThemeAndUser := `(SELECT q.id, q.question_text, q.max_mark FROM question AS q  WHERE q.theme=$1) EXCEPT (SELECT q.id, q.question_text, q.max_mark FROM question AS q JOIN question_answer AS qa ON qa.question_id=q.id WHERE q.theme=$1 AND qa.user_id=$2)`

	id := ctx.Value(middleware.CookieName)
	UUID, _ := id.(uuid.UUID)

	rows, err := r.db.Query(queryAllAnswersByThemeAndUser, theme, UUID)
	if err != nil {
		//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	questions := []*models.QuestionResp{}

	for rows.Next() {
		question := &models.QuestionResp{}
		err := rows.Scan(&question.ID, &question.QuestionText, &question.MaxMark)

		if err != nil {
			//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		questions = append(questions, question)
	}
	if err := rows.Err(); err != nil {
		//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	//utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return questions, nil
}

// SelectQuestionsByTheme retrieves question from the database by theme.
func (r *QuestionRepo) SelectAnswerStatistics(ctx context.Context) ([]*models.ThemeStatistic, error) {
	queryAllAnswersByThemeAndUser := `SELECT q.question_text, mark, COUNT(*) as answer_count
	FROM question_answer JOIN question as q ON q.id=question_answer.question_id GROUP BY q.question_text, mark`

	rows, err := r.db.Query(queryAllAnswersByThemeAndUser)
	if err != nil {
		//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	statistic := []*models.ThemeStatistic{}
	/*
		for rows.Next() {
			question := &models.QuestionResp{}
			err := rows.Scan(&question.ID, &question.QuestionText, &question.MaxMark)

			if err != nil {
				//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			questions = append(questions, question)
		}
		if err := rows.Err(); err != nil {
			//utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		//utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	*/
	return statistic, nil
}
