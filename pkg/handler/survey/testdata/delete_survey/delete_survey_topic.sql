INSERT INTO "public"."feedback_events" ("id", "deleted_at", "created_at", "updated_at", "title", "type", "subtype", "status", "created_by", "start_date", "end_date") VALUES
('8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', NULL, '2022-11-29 08:03:33.233262', '2022-11-29 08:03:33.233262', 'Q3/Q4, 2022', 'survey', 'peer-review', 'in-progress', '2655832e-f009-4b73-a535-64c3a22e558f', '2022-11-29 08:03:33.233262', '2023-05-29 08:03:33.233262');

INSERT INTO "public"."employee_event_topics" ("id", "deleted_at", "created_at", "updated_at", "title", "event_id", "employee_id", "project_id") VALUES
('e4a33adc-2495-43cf-b816-32feb8d5250e', NULL, '2022-11-29 08:03:33.233262', '2022-11-29 08:03:33.233262', 'Review Thuong Phung', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', '37e00d47-de69-4ac8-991b-cf3e39565c00', NULL);

INSERT INTO "public"."employee_event_reviewers" ("id", "deleted_at", "created_at", "updated_at", "event_id", "employee_event_topic_id", "reviewer_id", "relationship", "is_shared", "is_read", "author_status", "reviewer_status") VALUES
('e96999f5-b3f9-420c-9d9f-aa64e3e889ba', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'ecea9d15-05ba-4a4e-9787-54210e3b98ce', 'line-manager', 'f', 'f', 'draft', 'none'),
('bc9a5715-9723-4a2f-ad42-0d0f19a80b4a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '2655832e-f009-4b73-a535-64c3a22e558f', 'peer', 'f', 'f', 'draft', 'new'),
('c994db17-384a-4331-8944-b8ac0070ac3a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'd389d35e-c548-42cf-9f29-2a599969a8f2', 'peer', 'f', 'f', 'draft', 'new'),
('a7f4bcc2-1953-4b86-8aff-8a96f741fecb', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'f7c6016b-85b5-47f7-8027-23c2db482197', 'peer', 'f', 'f', 'draft', 'new'),
('16e87fe4-aff8-4c00-9cda-0b0efea94ada', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'd42a6fca-d3b8-4a48-80f7-a95772abda56', 'peer', 'f', 'f', 'draft', 'new'),
('cf850c59-24a2-45c5-98dc-bdf1818d36ba', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'dcfee24b-306d-4609-9c24-a4021639a11b', 'peer', 'f', 'f', 'draft', 'new'),
('f754ae24-8a82-4a8b-b93f-90bb38afa4da', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '38a00d4a-bc45-41de-965a-adc674ab82c9', 'peer', 'f', 'f', 'draft', 'new'),
('61a668bf-ba36-48e3-adc8-ce93c8aec29a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '608ea227-45a5-4c8a-af43-6c7280d96340', 'peer', 'f', 'f', 'draft', 'new'),
('c4b2c917-1912-40ee-8de2-e2dc1302a61a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '3f705527-0455-4e67-a585-6c1f23726fff', 'peer', 'f', 'f', 'draft', 'new'),
('77f64b43-93ed-48db-b20b-a3b2f030388a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'fae443f8-e8ff-4eec-b86c-98216d7662d8', 'peer', 'f', 'f', 'draft', 'none'),
('4f5ac428-9c23-4c5c-828f-bd9c77e5b56a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'ac318f73-0c8e-43ed-b00e-d230670dc400', 'peer', 'f', 'f', 'draft', 'none'),
('4d9b9f10-faee-42d6-ace3-99d0e96a0bea', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'a1f25e3e-cf40-4d97-a4d5-c27ee566b8c5', 'peer', 'f', 'f', 'draft', 'none'),
('6a7a936f-9ac8-4629-be67-0cf4c14f910a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '498d5805-dd64-4643-902d-95067d6e5ab5', 'peer', 'f', 'f', 'draft', 'none'),
('00e6a708-287c-4fc6-ae8d-83fed9d4263a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'f6ce0d0f-5794-463b-ad0b-8240ab9c49be', 'peer', 'f', 'f', 'draft', 'none'),
('68867dea-f8af-47e0-9bd7-620c0404f70a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '061820c0-bf6c-4b4a-9753-875f75d71a2c', 'peer', 'f', 'f', 'draft', 'none'),
('eae233c5-4930-4185-bfc7-7852f7b29fda', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'd675dfc5-acbe-4566-acde-f7cb132c0206', 'peer', 'f', 'f', 'draft', 'none'),
('4eecf200-4cd2-4df5-ae61-d4184613c0ba', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '8d7c99c0-3253-4286-93a9-e7554cb327ef', 'peer', 'f', 'f', 'draft', 'none'),
('fcdb0411-c790-4925-88ce-6b75c87d464a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'eeae589a-94e3-49ac-a94c-fcfb084152b2', 'peer', 'f', 'f', 'draft', 'none'),
('922e6243-c1e7-4055-9fe6-148ea39cfc8a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', 'd1b1dcbe-f4ed-49cc-a16b-56c9a3145c14', 'peer', 'f', 'f', 'draft', 'none'),
('359a2804-40a8-4a6d-86fb-c0247f66c28a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '7bcf4b45-0279-4da2-84e4-eec5d9d05ba3', 'peer', 'f', 'f', 'draft', 'none'),
('d896ce40-dd8c-48be-951e-7691aa27a73a', NULL, '2022-12-05 16:26:15.411511', '2022-12-05 16:26:15.411511', '8a5bfedb-6e11-4f5c-82d9-2635cfcce3e1', 'e4a33adc-2495-43cf-b816-32feb8d5250e', '7fbfb59b-e00e-46b2-85cd-64f9f9942daa', 'peer', 'f', 'f', 'draft', 'none');


