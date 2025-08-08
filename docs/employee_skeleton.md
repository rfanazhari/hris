# Domain Model: Employee

## Employee
- `id` : UUID
- `personalInfo` : PersonalInfo (Value Object)
- `contactInfos` : ContactInfo[] (Value Object)
- `jobPositions` : JobPosition[] (Entity / Value Object tergantung kebutuhan)
- `employmentContracts` : EmploymentContract[] (Entity)
- `documents` : Document[] (Value Object)
- `salaryRecords` : SalaryRecord[] (Entity)
- `emergencyContacts` : EmergencyContact[] (Value Object)
- `status` : EmployeeStatus (Enum)
- `createdAt` : DateTime
- `updatedAt` : DateTime

---

## PersonalInfo (Value Object)
- `fullName` : string
- `birthDate` : Date
- `gender` : Gender (Enum)
- `nationality` : string
- `maritalStatus` : MaritalStatus (Enum)
- `religion` : string (optional)
- `placeOfBirth` : string

---

## ContactInfo (Value Object)
- `type` : ContactType (Enum: PRIMARY, SECONDARY, EMERGENCY)
- `phone` : PhoneNumber (Value Object)
- `email` : EmailAddress (Value Object)
- `address` : Address (Value Object)

---

## JobPosition
- `id` : UUID
- `title` : string
- `gradeLevel` : GradeLevel (Enum)
- `department` : Department (Entity atau Value Object)
- `startDate` : Date
- `endDate` : Date? (nullable)

---

## EmploymentContract (Entity)
- `id` : UUID
- `contractType` : ContractType (Enum: PKWT, PKWTT, Freelance, Internship)
- `startDate` : Date
- `endDate` : Date? (nullable)
- `document` : Document (Value Object)
- `status` : ContractStatus (Enum: Active, Expired, Terminated)

---

## Document (Value Object)
- `type` : DocumentType (Enum)
- `file` : FileReference (Value Object)
- `validity` : ValidityPeriod (Value Object)

---

## SalaryRecord (Entity)
- `id` : UUID
- `amount` : number
- `currency` : string
- `effectiveDate` : Date
- `bonus` : number (optional)

---

## EmergencyContact (Value Object)
- `name` : string
- `relationship` : string
- `phone` : PhoneNumber (Value Object)
- `email` : EmailAddress (optional)

---

## EmployeeStatus (Enum)
- ACTIVE
- INACTIVE
- ON_LEAVE
- RESIGNED
- TERMINATED

---

## Note:
- Semua tanggal menggunakan format ISO 8601 (YYYY-MM-DD).
- UUID adalah string dengan format UUID v4.
- Value Object artinya immutable dan tidak punya identity sendiri, hanya sebagai atribut.
- Entity memiliki identity (biasanya UUID) dan bisa berubah status.
- Enum adalah nilai tetap yang sudah didefinisikan dan tidak bisa sembarangan diisi.  
