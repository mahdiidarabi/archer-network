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

	fmt.Println("Message is received by Archer ", a.ArcherID, "at time ", time.Now())

	duration := a.TotalNumberOfArchers - a.ArcherID

	switch msgStr {
	case fireOneMsg:
		{
			if duration == 1 {
				// the last archer
				go a.Fire(duration)

				return nil

			} else {

				go a.Fire(duration)

				err := a.BroadcastMessageToNeighbors(ctx, msgStr)
				if err != nil {
					return err
				}

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

	oneSecond := time.Second
	time.Sleep(oneSecond)

	err := a.RightNeighbor.HearFromNeighborsForMessage(ctx, msgStr)
	if err != nil {
		return err
	}

	return nil
}

func (a *archer) Fire(duration uint64) {
	durationInSecond := time.Duration(duration)*time.Second
	time.Sleep(durationInSecond)

	fmt.Println("this is the fire message from archer ", a.ArcherID, "and the time is ", time.Now())
}

func (a *archer) AddLeftNeighbor(leftNeighbor Archer)  {
	a.LeftNeighbor = leftNeighbor
}

func (a *archer) AddRightNeighbor(rightNeighbor Archer)  {
	a.RightNeighbor = rightNeighbor
}
