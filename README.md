Convert Decimal Number to Thai Baht Text (Go)

This project is a Go implementation for converting decimal numbers into Thai text
representation with Baht (บาท) and Satang (สตางค์) currency suffixes.

It uses the github.com/shopspring/decimal package to ensure precision-safe
decimal handling and demonstrates how such logic can be integrated into backend
services.

======================================================================
Problem Description
======================================================================

Convert a decimal value into Thai text with currency formatting.

Rules:
- The input is a decimal value
- The output must be a Thai text string
- If there is no fractional part, append "บาทถ้วน"
- If there is a fractional part, convert it to Satang and append "สตางค์"
- Fractional part is handled up to 2 decimal places

======================================================================
Examples
======================================================================

Input: 0
Output: ศูนย์บาทถ้วน

Input: 100
Output: หนึ่งร้อยบาทถ้วน

Input: 101
Output: หนึ่งร้อยหนึ่งบาทถ้วน

Input: 110
Output: หนึ่งร้อยสิบบาทถ้วน

Input: 1234
Output: หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน

Input: 33333.75
Output: สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์

======================================================================
Project Structure
======================================================================

.
├── go.mod
├── go.sum
├── main.go
└── README.txt

======================================================================
Requirements
======================================================================

- Go 1.20 or higher
- Dependency:
  - github.com/shopspring/decimal

======================================================================
Installation
======================================================================

Clone the repository:

git clone https://github.com/anikSomjitchob/convertNumberToTHB-text.git
cd convertNumberToTHB-text

Install dependencies:

go mod tidy

======================================================================
Running the Program
======================================================================

Run the application using:

go run .

The program prints Thai Baht text output for predefined decimal inputs
configured in main.go.

======================================================================
Code Overview
======================================================================

main()
- Defines example decimal inputs
- Calls the conversion function
- Prints Thai Baht text results

covertNumbersToTH(input decimal.Decimal)
- Converts decimal input into Thai Baht text
- Handles integer and fractional parts
- Appends correct currency suffix

assignNumberWord(number, location, max)
- Converts numeric digits to Thai number words
- Applies Thai language rules such as:
  - ยี่ for twenty
  - Omitting หนึ่ง before สิบ
  - Correct zero handling

assignLocationWord(number, location)
- Maps digit position to Thai units:
  - สิบ, ร้อย, พัน, หมื่น, แสน, ล้าน

======================================================================
Service Integration Considerations
======================================================================

Although implemented as a console application, the conversion logic can be
extracted into a reusable package and integrated into:

- REST APIs
- gRPC services
- Financial systems (invoices, receipts, reports)

Example usage in a service:

result := covertNumbersToTH(amount)

======================================================================
Notes
======================================================================

- Fractional values are truncated to 2 decimal places
- The logic is stateless and safe for concurrent usage
- Designed as a foundation that can be refactored into a shared utility package

======================================================================
Author
======================================================================

Anik Somjitchob
