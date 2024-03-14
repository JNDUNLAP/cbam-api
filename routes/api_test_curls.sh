# #!/bin/bash

# # Base URL
# BASE_URL="http://localhost:1323"

# # Ping test
# echo "Ping Test"
# curl "${BASE_URL}/api/ping"

# # Quarterly Reports
# echo "\n\nGet Quarterly Reports"
# curl "${BASE_URL}/quarterly-reports"

# echo "\n\nGet Specific Quarterly Report"
# curl "${BASE_URL}/quarterly-reports/1"

# echo "\n\nCreate Quarterly Report"
# curl -X POST "${BASE_URL}/quarterly-reports" -H "Content-Type: application/json" -d '{}'

# echo "\n\nUpdate Quarterly Report"
# curl -X PUT "${BASE_URL}/quarterly-reports/1" -H "Content-Type: application/json" -d '{}'

# echo "\n\nDelete Quarterly Report"
# curl -X DELETE "${BASE_URL}/quarterly-reports/1"

# # Declarant
# echo "\n\nGet Declarant"
# curl "${BASE_URL}/quarterly-reports/1/declarant"

# echo "\n\nUpdate Declarant"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/declarant" -H "Content-Type: application/json" -d '{}'

# # Importer
# echo "\n\nGet Importer"
# curl "${BASE_URL}/quarterly-reports/1/importer"

# echo "\n\nUpdate Importer"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/importer" -H "Content-Type: application/json" -d '{}'

# # National Competent Auth and Signatures
# echo "\n\nGet National Competent Auth"
# curl "${BASE_URL}/quarterly-reports/1/national-competent-auth"

# echo "\n\nUpdate National Competent Auth"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/national-competent-auth" -H "Content-Type: application/json" -d '{}'

# echo "\n\nGet Signatures"
# curl "${BASE_URL}/quarterly-reports/1/signatures"

# echo "\n\nUpdate Signatures"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/signatures" -H "Content-Type: application/json" -d '{}'

# # Imported Goods
# echo "\n\nGet Imported Goods"
# curl "${BASE_URL}/quarterly-reports/1/imported-goods"

# echo "\n\nAdd Imported Good"
# curl -X POST "${BASE_URL}/quarterly-reports/1/imported-goods" -H "Content-Type: application/json" -d '{}'

# echo "\n\nUpdate Imported Good"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/imported-goods/100" -H "Content-Type: application/json" -d '{}'

# echo "\n\nDelete Imported Good"
# curl -X DELETE "${BASE_URL}/quarterly-reports/1/imported-goods/100"

# # Goods Emissions
# echo "\n\nGet Goods Emissions"
# curl "${BASE_URL}/quarterly-reports/1/imported-goods/100/emissions"

# echo "\n\nUpdate Goods Emissions"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/imported-goods/100/emissions" -H "Content-Type: application/json" -d '{}'

# # Supporting Documents
# echo "\n\nGet Supporting Documents"
# curl "${BASE_URL}/quarterly-reports/1/imported-goods/100/supporting-documents"

# echo "\n\nAdd Supporting Document"
# curl -X POST "${BASE_URL}/quarterly-reports/1/imported-goods/100/supporting-documents" -H "Content-Type: application/json" -d '{}'

# echo "\n\nGet Specific Supporting Document"
# curl "${BASE_URL}/quarterly-reports/1/imported-goods/100/supporting-documents/500"

# echo "\n\nUpdate Supporting Document"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/imported-goods/100/supporting-documents/500" -H "Content-Type: application/json" -d '{}'

# echo "\n\nDelete Supporting Document"
# curl -X DELETE "${BASE_URL}/quarterly-reports/1/imported-goods/100/supporting-documents/500"

# # Remarks Emissions
# echo "\n\nGet Remarks Emissions"
# curl "${BASE_URL}/quarterly-reports/1/imported-goods/100/remarks-emissions"

# echo "\n\nUpdate Remarks Emissions"
# curl -X PUT "${BASE_URL}/quarterly-reports/1/imported-goods/100/remarks-emissions" -H "Content-Type: application/json" -d '{}'

# echo "\nAll API tests completed."
