package aws_client

import (
	"context"
	"golang.org/x/sync/semaphore"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type ListResolverFunc func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, detailChan chan<- any) error

type DetailResolverFunc func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask,
	resultChannel chan<- any, errorChan chan<- error, summary interface{})

func ListAndDetailResolver(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any,
	list ListResolverFunc, details DetailResolverFunc) error {
	errorChan := make(chan error)
	detailChan := make(chan interface{})
	c := client.(*Client)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for detailError := range errorChan {

			clientMeta.WarnF("Error while fetching details: %s", detailError.Error())
		}
	}()
	sem := semaphore.NewWeighted(int64(MAX_GOROUTINES))

	go func() {
		defer close(errorChan)
		for item := range detailChan {
			if err := sem.Acquire(ctx, 1); err != nil {
				continue
			}
			func(summary interface{}) {
				defer sem.Release(1)

				details(ctx, clientMeta, c, task, resultChannel, errorChan, summary)
			}(item)
		}
	}()

	err := list(ctx, clientMeta, client, task, detailChan)
	close(detailChan)
	if err != nil {
		return err
	}

	<-done

	return nil
}
