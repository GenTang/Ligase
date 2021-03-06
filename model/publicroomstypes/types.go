// Copyright 2017 Vector Creations Ltd
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
//
// Modifications copyright (C) 2020 Finogeeks Co., Ltd

package publicroomstypes

// PublicRoom represents a local public room
type PublicRoom struct {
	RoomID           string   `json:"room_id"`
	Aliases          []string `json:"aliases,omitempty"`
	CanonicalAlias   string   `json:"canonical_alias,omitempty"`
	Name             string   `json:"name,omitempty"`
	Topic            string   `json:"topic,omitempty"`
	AvatarURL        string   `json:"avatar_url,omitempty"`
	NumJoinedMembers int64    `json:"num_joined_members"`
	WorldReadable    bool     `json:"world_readable"`
	GuestCanJoin     bool     `json:"guest_can_join"`
}
