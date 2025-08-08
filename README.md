# HRIS Backend Domain Model (Go)

## Overview
This repository contains a backend domain model for an HRIS (Human Resources Information System) implemented in Go. The project focuses on clear domain boundaries, strong typing, and testable business rules using Domain-Driven Design (DDD) concepts: Entities, Value Objects, Enums, and Factories.

The primary objective is to model core HR concepts (employee personal information, job positions, organization units, contact and document handling, salary ranges, and statuses) in a clean and validated way, serving as a foundation for future application layers (APIs, persistence, messaging). The current codebase emphasizes correctness and expressiveness of the domain rather than runtime features.

## Goals
- Core HR (Personnel Administration): model employee master data, job/position management, organizational structure, document management, contract history, and compliance (BPJS, tax, labor regulations). 
- Talent Acquisition: represent recruitment flows (job posting, ATS, screening), new-hire onboarding, and internal mobility (transfer/promotion) as domain processes and states.
- Time & Attendance: capture attendance, shift scheduling, leave/permit, and overtime policies with clear rules and validations.
- Payroll & Compensation: express payroll and benefits, deductions (tax, BPJS, loan), bonuses & incentives, and payslip artifacts with precise value objects.
- Performance Management: support KPI/OKR tracking, periodic/annual reviews, 360° feedback, and career pathing structures.
- Learning & Development: cover training management, certifications & skill tracking, and e-learning references.
- ESS & MSS: enable employee self-service for personal data access, leave/overtime/reimbursement requests, payslip/schedule/document access, and manager approval workflows.
- Analytics & Reporting: provide domain outputs to support HR reports (turnover, attendance, payroll), predictive analytics (e.g., attrition), and executive dashboards.

## Key Features
- Domain entities with private fields and read-only getters to protect invariants.
- Value Objects for names, addresses, phone numbers, emails, salary ranges, document validity, and more.
- Enumerations for constrained domains (gender, nationality, marital status, religion, contact/document types, grade levels, etc.) with parsing and validation support.
- Factory types that centralize construction rules and validations for entities (e.g., JobPositionFactory, PersonalInfoFactory).
- Comprehensive unit tests covering enums, value objects, and factories.
- Utility packages for validation helpers and fake data generation.


## Design Highlights
- Domain-Driven Design: Entities expose intent via behavior/getters; construction is guarded by factories to ensure valid state from inception.
- Validation:
  - Enums provide ParseXxx helpers with strict allowed values and case-insensitive input handling.
  - Value Objects normalize and validate inputs (e.g., EmployeeName trims and requires first/last names; SalaryRange enforces min/max and currency correctness).
  - Factories apply business rules (e.g., minimum title/description lengths in JobPositionFactory; non-empty place of birth, valid enumerations in PersonalInfoFactory).
- Immutability: Value Objects are designed to be immutable after creation.
- Testability: Each enum/value object/factory has dedicated tests to ensure domain rules remain stable.

## Getting Started
Prerequisites:
- Go 1.24 or later (as specified in go.mod).

Common tasks:
- Run all tests: go test ./...
- Work on domain code: modify or extend types under domain/entity, domain/valueobject, and domain/enum. Add or update tests accordingly.

Note: The main package (cmd/main.go) is intentionally minimal; this project currently serves as a domain module. You can integrate it into a service by adding application/infrastructure layers (HTTP handlers, persistence, etc.).

## Documentation
- See docs/employee_skeleton.md for a conceptual outline of the Employee aggregate and related components.

## Roadmap
- Phase 1 — Core HR (Personnel Administration): implement Employee aggregate root (PersonalInfo, ContactInfo, EmploymentContract, Document, OrganizationUnit, JobPosition), repositories and basic domain services; enforce validations and add unit tests.
- Phase 2 — Time & Attendance: model attendance records, shift scheduling, leave/permit, and overtime policies with factories/value objects and comprehensive tests.
- Phase 3 — Payroll & Compensation: represent salary structures, benefits/deductions (tax, BPJS, loans), bonuses & incentives, and payslip artifacts; prepare calculation scaffolding and validations.
- Phase 4 — Talent Acquisition: define job posting and ATS pipeline states, screening processes, and onboarding integration points as domain processes and states.
- Phase 5 — Performance Management: support KPI/OKR models, review cycles (periodic/annual), 360° feedback entities, and career pathing structures.
- Phase 6 — Learning & Development: training management, certifications tracking, skills catalog, and references to e-learning.
- Phase 7 — ESS & MSS: self-service requests (personal data access, leave/overtime/reimbursement), payslip/schedule/document access, and manager approval workflows.
- Phase 8 — Analytics & Reporting: domain outputs/aggregates to support HR reports and executive dashboards; prepare hooks for predictive analytics.
- Cross-cutting (incremental across phases): introduce application layer (use cases), REST/GraphQL endpoints, and persistence mappings/migrations as the domain stabilizes.

## License
:(