API_TOKEN=$(cat /etc/secrets/api-token)
FROM_EMAIL=$(cat /etc/secrets/from-email)

curl -X POST "https://api.mailersend.com/v1/email" \
    -H "Authorization: Bearer $API_TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
          "from": {
            "email": "'"$FROM_EMAIL"'"
          },
          "to": [
            {
              "email": "candyahs@gmail.com"
            }
          ],
          "subject": "Test Email from Temporary Pod",
          "text": "This is a test email from the temporary pod.",
          "html": "<p>This is a test email from the temporary pod.</p>"
        }'
