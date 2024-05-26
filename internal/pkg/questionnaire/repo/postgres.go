package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"context"
	"database/sql"
	"fmt"

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
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// SelectQuestionsByTheme retrieves question from the database by theme.
func (r *QuestionRepo) SelectQuestionsByTheme(ctx context.Context, theme *models.QuestionTheme) ([]*models.QuestionResp, error) {
	queryAllAnswersByThemeAndUser := `(SELECT q.id, q.question_text, q.max_mark FROM question AS q  WHERE q.theme=$1) EXCEPT (SELECT q.id, q.question_text, q.max_mark FROM question AS q JOIN question_answer AS qa ON qa.question_id=q.id WHERE q.theme=$1 AND qa.user_id=$2)`

	id := ctx.Value(middleware.CookieName)
	IdUser, _ := id.(int64)

	rows, err := r.db.Query(queryAllAnswersByThemeAndUser, *theme, IdUser)

	if err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	questions := []*models.QuestionResp{}

	for rows.Next() {
		question := &models.QuestionResp{}
		err := rows.Scan(&question.ID, &question.QuestionText, &question.MaxMark)

		if err != nil {
			// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		questions = append(questions, question)
	}
	if err := rows.Err(); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return questions, nil
}

func (r *QuestionRepo) SelectAnswerStatistics(ctx context.Context) ([]*models.ThemeStatistic, error) {
	query := `
        SELECT q.question_text, qa.mark, q.theme, COUNT(*) as answer_count
        FROM question_answer qa
        JOIN question q ON q.id = qa.question_id
        GROUP BY q.question_text, qa.mark, q.theme
        ORDER BY q.theme, q.question_text, qa.mark
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}
	defer rows.Close()

	var themesMap = make(map[models.QuestionTheme]*models.ThemeStatistic)
	for rows.Next() {
		var questionText, theme string
		var mark, answerCount int
		if err := rows.Scan(&questionText, &mark, &theme, &answerCount); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		themeStat, ok := themesMap[models.QuestionTheme(theme)]
		if !ok {
			themeStat = &models.ThemeStatistic{Theme: models.QuestionTheme(theme), Questions: []*models.QuestionAnswerStatisticsResp{}}
			themesMap[models.QuestionTheme(theme)] = themeStat
		}

		questionStat := findQuestionAnswerStatisticsResp(themeStat.Questions, questionText)
		if questionStat == nil {
			questionStat = &models.QuestionAnswerStatisticsResp{Title: questionText, QuestionsTopic: []*models.QuestionWithTitleStat{}}
			themeStat.Questions = append(themeStat.Questions, questionStat)
		}

		questionStat.QuestionsTopic = append(questionStat.QuestionsTopic, &models.QuestionWithTitleStat{CountAnswers: answerCount, Mark: mark})
	}

	var themes []*models.ThemeStatistic
	for _, v := range themesMap {
		themes = append(themes, v)
	}

	return themes, nil
}

func findQuestionAnswerStatisticsResp(questions []*models.QuestionAnswerStatisticsResp, title string) *models.QuestionAnswerStatisticsResp {
	for _, q := range questions {
		if q.Title == title {
			return q
		}
	}
	return nil
}
