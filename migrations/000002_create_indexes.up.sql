CREATE INDEX IF NOT EXISTS idx_user_team_id ON users (team_id);
CREATE INDEX IF NOT EXISTS idx_pr_author_id ON pull_requests (author_id);
CREATE INDEX IF NOT EXISTS idx_reviewers_user_id ON reviewers (user_id);