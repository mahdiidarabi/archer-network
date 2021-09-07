package domain

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	fireOneMsg = "fire-1"
)

type Archer interface {
	HearFromNeighborsForMessage(context.Context, string) error
	BroadcastMessageToNeighbors(context.Context, string) error
	Fire(uint64)
	AddLeftNeighbor(Archer)
	AddRightNeighbor(Archer)
}

type archer struct {
	TotalNumberOfArchers uint64
	ArcherID uint64
	LeftNeighbor Archer
	RightNeighbor Archer
}

func NewArcher(totalNum uint64, id uint64) Archer {
	return &archer{
		TotalNumberOfArchers: totalNum,
		ArcherID:             id,
	}
}

func (a *archer) HearFromNeighborsForMessage(ctx context.Context, msgStr string) error {

	duration := a.TotalNumberOfArchers - a.ArcherID

	switch msgStr {
	case fireOneMsg:
		{
			if duration == 1 {
				// the last archer
				a.Fire(duration)

				return nil

			} else {
				err := a.BroadcastMessageToNeighbors(ctx, msgStr)
				if err != nil {
					return err
				}

				a.Fire(duration)
				return nil
			}
		}
	default:
		{
			return errors.New("this is not a known message")
		}
	}
}

func (a *archer) BroadcastMessageToNeighbors(ctx context.Context, msgStr string) error {

	if a.RightNeighbor == nil {
		return errors.New("there is no right archer to broadcast the message to it")
	}

	err := a.RightNeighbor.HearFromNeighborsForMessage(ctx, msgStr)
	if err != nil {
		return err
	}

	oneSecond := time.Second
	time.Sleep(oneSecond)

	return nil
}

func (a *archer) Fire(duration uint64) {
	durationInSecond := time.Duration(duration)*time.Second
	time.Sleep(durationInSecond)

	fmt.Println("this is the fire message from archer ", a.ArcherID)
	fmt.Println(time.Now())
}

func (a *archer) AddLeftNeighbor(leftNeighbor Archer)  {
	a.LeftNeighbor = leftNeighbor
}

func (a *archer) AddRightNeighbor(rightNeighbor Archer)  {
	a.RightNeighbor = rightNeighbor
}
