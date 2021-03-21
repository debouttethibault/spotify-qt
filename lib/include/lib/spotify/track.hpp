#pragma once

#include "thirdparty/json.hpp"
#include "lib/format.hpp"

#include <string>

namespace lib
{
	namespace spt
	{
		/**
		 * A music track
		 */
		class track
		{
		public:
			track() = default;

			/**
			 * Local track, not streamed from Spotify
			 */
			bool is_local = false;

			/**
			 * Playable in current region
			 */
			bool is_playable = false;

			/**
			 * Track duration in milliseconds
			 */
			int duration = 0;

			/**
			 * ISO date when it was added to the context
			 */
			std::string added_at;

			/**
			 * Name of album
			 */
			std::string album = "(no album)";

			/**
			 * ID of album
			 */
			std::string album_id = "0";

			/**
			 * Name of artist
			 * @todo Add support for multiple artists
			 */
			std::string artist = "(no artist)";

			/**
			 * ID of artist
			 * @todo Add support for multiple artists
			 */
			std::string artist_id = "0";

			/**
			 * ID of track
			 */
			std::string id = "0";

			/**
			 * URL to cover of album
			 */
			std::string image;

			/**
			 * Name of track
			 */
			std::string name = "(no name)";

			/**
			 * Format track as "{name} - {artist}"
			 */
			std::string title();
		};

		/**
		 * track -> json
		 */
		void to_json(nlohmann::json &j, const track &t);

		/**
		 * json -> track
		 */
		void from_json(const nlohmann::json &j, track &t);
	}
}
