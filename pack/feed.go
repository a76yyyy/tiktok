// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"github.com/a76yyyy/tiktok/kitex_gen/feed"

	"github.com/a76yyyy/tiktok/dal/db"
)

// Video pack feed info
func Video(u *db.Video) *feed.Video {
	if u == nil {
		return nil
	}

	author := User(&u.Author)
	favorite_count := int64(u.FavoriteCount)
	comment_count := int64(u.CommentCount)

	return &feed.Video{
		Id:            int64(u.ID),
		Author:        author,
		PlayUrl:       u.PlayUrl,
		CoverUrl:      u.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		Title:         u.Title,
	}
}

// Videos pack list of user info
func Videos(us []*db.Video) []*feed.Video {
	videos := make([]*feed.Video, 0)
	for _, u := range us {
		if user2 := Video(u); user2 != nil {
			videos = append(videos, user2)
		}
	}
	return videos
}
