//go:generate wire
//go:build wireinject

package main

import (
	"github.com/asynccnu/ccnubox-be/be-grade/cron"
	"github.com/asynccnu/ccnubox-be/be-grade/events"
	"github.com/asynccnu/ccnubox-be/be-grade/events/producer"
	"github.com/asynccnu/ccnubox-be/be-grade/grpc"
	"github.com/asynccnu/ccnubox-be/be-grade/ioc"
	"github.com/asynccnu/ccnubox-be/be-grade/repository/dao"
	"github.com/asynccnu/ccnubox-be/be-grade/service"
	"github.com/google/wire"
)

func InitApp() App {
	wire.Build(
		events.NewGradeDetailEventConsumerHandler,
		producer.NewSaramaProducer,
		grpc.NewGradeGrpcService,
		service.NewGradeService,
		dao.NewGradeDAO,
		// 第三方
		ioc.InitEtcdClient,
		ioc.InitDB,
		ioc.InitLogger,
		ioc.InitGRPCxKratosServer,
		ioc.InitUserClient,
		ioc.InitCounterClient,
		ioc.InitFeedClient,
		ioc.InitClasslistClient,
		ioc.InitKafka,
		ioc.InitConsumers,
		cron.NewGradeController,
		cron.NewCron,
		NewApp,
	)
	return App{}
}
