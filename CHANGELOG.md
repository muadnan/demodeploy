# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [Unreleased]

### Added

- Verify user and group existence before creating/updating role assignments (CI-503).
- Invalidate access object cache when group membership or role assignments are changed (CI-465). 
- Tenant Provision API in authz service to add role assignments for admin user (CI-598).
- Grafana Dashboards (CI-583).
- Enabled TLS for redis connection (CI-592).
- Implement templating for seeding data in authz service (CI-610).
- Script for making all users in a tenant admin (CI-628).
- Script to added changelog for renovate (SRE-152).

### Changed

- [BREAKING] Make Contexts type safe (CI-529).
- Only load dataloaders that are required in the graphql request (CI-537).
- [BREAKING] Updated user and group filtering in keycloak (CI-584).
