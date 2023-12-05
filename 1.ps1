$uri = 'https://api-free.deepl.com/v2/translate'
$headers = @{
            "Authorization" = "DeepL-Auth-Key "
}
$body = @{
            "text" = "Join Senior Developer Advocate Christina Warren for all the latest developer. open source, and GitHub news.
            hello.
            liwei.
            "
                "source_lang" = "EN"
                    "target_lang" = "ZH"
}

$response = Invoke-RestMethod -Uri $uri -Method 'Post' -Headers $headers -Body $body -ContentType 'application/x-www-form-urlencoded'

Write-Output $response.translations[0].text
