package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeUnlinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) (*api.AccountApple, error) {
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RegisterBeforeUnlinkApple", "in": in}).Debug("")
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	_, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkApple", "in": in}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return in, err
	// }
	return in, nil
}
