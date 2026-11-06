# Business events V1

`EventV1` is the shared journal envelope. `DecodeV1` validates the envelope once at the consumer boundary. Data uses `json.RawMessage` to retain SQL nulls and exact numbers.

The `postgres-offer-publication-trigger` sender requires one of the three offer actions, a consistent offer identity, a nonempty sorted unique root `changed_fields` array, and exactly the flat `OfferSnapshotColumns` keys. Snapshot values remain stored database values, including nulls and PostgreSQL timestamp infinity strings. Large text values never enter the snapshot. `OfferSiteID` handles missing, null, non-integer, or nonpositive routing keys separately from decoding.

Retained application envelopes can omit `changed_fields` and `object_id`, and retain their original data shape. Consumers select the sender before acting on historical API content events. `testdata/offer-published.json` is the language-neutral fixture for migration and consumer tests.

Rollout preserves the existing domain producer and subscription. Install a dual-path worker, remove every legacy API publication/content writer, then activate the forward trigger migration. The API-only intermediate state still uses domain status events. The trigger must not run alongside old API writers. Retire the frontend domain subscription only after business-path acceptance. Source workspace tests do not authorize publishing or activation. Consumers need a published module pin and `GOWORK=off` checks before release.
