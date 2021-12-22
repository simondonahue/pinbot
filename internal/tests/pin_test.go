package tests

import (
	"testing"
)

func TestPin(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		the_message_is_posted()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		a_pin_message_should_be_posted_in_the_last_channel().and().
		the_bot_should_add_the_emoji("👀").and().
		the_bot_should_add_the_emoji("✅")
}

func TestPinGeneralPinsChannel(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		a_channel_named("pins").and().
		the_message_is_posted()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		a_pin_message_should_be_posted_in_the_last_channel()
}

func TestPinSpecificPinsChannel(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		a_channel_named("pins").and().
		a_channel_named("test-pins").and().
		the_message_is_posted()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		a_pin_message_should_be_posted_in_the_last_channel()
}

func TestPinDuplicate(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		the_message_is_posted().and().
		the_message_is_already_pinned()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		the_bot_should_log_the_message_as_already_pinned()
}

func TestPinSelfPinDisabled(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		self_pin_is_disabled().and().
		the_message_is_posted().and().
		the_message_is_already_pinned()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		the_bot_should_log_the_message_as_an_avoided_self_pin()
}

func TestPinImportCommand(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		the_message_is_posted().and().
		the_message_is_pinned()

	when.
		an_import_is_triggered()

	then.
		a_pin_message_should_be_posted_in_the_last_channel().and().
		the_bot_should_add_the_emoji("👀").and().
		the_bot_should_add_the_emoji("✅")
}

func TestPinImportCommandIgnoreAlreadyPinned(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		the_message_is_posted().and().
		the_message_is_pinned()

	when.
		an_import_is_triggered()

	then.
		a_pin_message_should_be_posted_in_the_last_channel().and().
		the_bot_should_add_the_emoji("👀").and().
		the_bot_should_add_the_emoji("✅")
}

func TestPinWithImage(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		a_message().and().
		an_image_attachment().and().
		the_message_is_posted()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		a_pin_message_should_be_posted_in_the_last_channel().and().
		the_pin_message_should_have_n_embeds(1).and().
		the_pin_message_should_have_an_image_embed()
}

func TestPinWithMultipleImage(t *testing.T) {
	given, when, then := NewPinStage(t)

	given.
		a_channel_named("test").and().
		a_message().and().
		an_image_attachment().and().
		another_image_attachment().and().
		the_message_is_posted()

	when.
		the_message_is_reacted_to_with("📌")

	then.
		a_pin_message_should_be_posted_in_the_last_channel().and().
		the_pin_message_should_have_n_embeds(3).and().
		the_pin_message_should_have_n_embeds_with_url(2)
}
