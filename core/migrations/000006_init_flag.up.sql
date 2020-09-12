BEGIN;

CREATE TABLE flag (
  id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  -- attributes
  key resource_key,
  name resource_name,
  description resource_description,
  tags resource_tags,
  -- references
  project_id UUID REFERENCES project (id),
  -- contraints
  CONSTRAINT flag_key UNIQUE(key, project_id)
);

CREATE TABLE variation (
  id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  -- attributes
  key resource_key,
  name resource_name,
  description resource_description,
  tags resource_tags,
  -- references
  flag_id UUID REFERENCES flag (id),
  -- contraints
  CONSTRAINT variation_key UNIQUE(key, flag_id)
);

END;
