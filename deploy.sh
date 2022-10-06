# EvidenceService function
gcloud functions deploy evidence-service \
--gen2 \
--runtime=go116 \
--region=us-west1 \
--source=. \
--entry-point=EvidenceService \
--trigger-http \
--allow-unauthenticated
