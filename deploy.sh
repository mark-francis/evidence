# EvidenceDownload function
gcloud functions deploy evidence-download \
--gen2 \
--runtime=go116 \
--region=us-west1 \
--source=. \
--entry-point=EvidenceDownload \
--trigger-http \
--allow-unauthenticated

# EvidenceUpload function
gcloud functions deploy evidence-upload \
--gen2 \
--runtime=go116 \
--region=us-west1 \
--source=. \
--entry-point=EvidenceUpload \
--trigger-http \
--allow-unauthenticated
