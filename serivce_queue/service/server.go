package service

import (
	"context"
	"service_queue/pb"
)

type JudgeService struct {
	job_queue []*pb.Submission
	pb.UnimplementedJudgeServiceServer
}

func NewJudgeService() *JudgeService {
	return &JudgeService{
		job_queue: make([]*pb.Submission, 5),
	}
}

func (s *JudgeService) Submit(ctx context.Context, submission *pb.Submission) (*pb.Empty, error) {
	s.job_queue = append(s.job_queue, submission)
	return &pb.Empty{}, nil
}

func (s *JudgeService) GetJob(ctx context.Context, worker *pb.Empty) (*pb.Job, error) {
	if len(s.job_queue) == 0 {
		return &pb.Job{ValidJob: false, Submission: nil}, nil
	} else {
		got_job := s.job_queue[0]
		s.job_queue = s.job_queue[1:]

		return &pb.Job{
			ValidJob:   true,
			Submission: got_job,
		}, nil
	}
}
