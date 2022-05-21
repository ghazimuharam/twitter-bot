package mock

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -source=../internal/twitter/usecase.go --build_flags=--mod=mod -destination=mock_usecase.go -package=mock
//go:generate mockgen -source=../internal/twitter/repository.go --build_flags=--mod=mod -destination=mock_repository.go -package=mock
//go:generate mockgen -source=../internal/twitter/repository/client/client.go --build_flags=--mod=mod -destination=mock_twitter_client.go -package=mock
