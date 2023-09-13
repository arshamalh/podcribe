package handlers

import (
	"net/http"
	"podcribe/api/requests"
	"podcribe/manager"
	"podcribe/services/convertor"
	"podcribe/services/crawler"
	"podcribe/services/downloader"
	"podcribe/services/transcriber"
	"podcribe/services/translator"

	"github.com/labstack/echo/v4"
)

func Process() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		body := new(requests.CreateWorkFlowAfterLink)
		if err := ctx.Bind(body); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		// TODO: Make a websocket connection to show the progress
		// using downloader.ConsumeProgress()
		downloader := downloader.New(3)

		manager := manager.New(
			crawler.New(), downloader, convertor.New(),
			transcriber.New(), translator.New(),
		)
		if body.IsJustDownload {
			filepath, err := manager.JustDownload(body.Link)
			// TODO: Turn filepath to a downloadable link
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, err)
			}
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"filepath": filepath,
			})
		}
		if body.IsNoTranslation {
			podcast, err := manager.FullExceptTranslation(body.Link)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, err)
			}
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"filepath":      podcast.Mp3Path,
				"transcription": podcast.TranscriptionPath,
			})
		}

		podcast, err := manager.FullFlow(body.Link)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		return ctx.JSON(http.StatusOK, echo.Map{
			"translation":   podcast.TranslationPath,
			"filepath":      podcast.Mp3Path,
			"transcription": podcast.TranscriptionPath,
		})
	}
}
