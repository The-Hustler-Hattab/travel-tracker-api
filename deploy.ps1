# Deploy the app to GCP Cloud Run
docker build -t travel-tracker-api:latest .

# Tag the Docker image
docker tag travel-tracker-api:latest gcr.io/personal-projects-416300/travel-tracker-api:latest

# Push the Docker image to Google Container Registry (GCR)
docker push gcr.io/personal-projects-416300/travel-tracker-api:latest

# Deploy the app to Cloud Run
gcloud run deploy travel-tracker-api `
  --image gcr.io/personal-projects-416300/travel-tracker-api:latest `
  --platform managed `
  --region us-central1 `
  --allow-unauthenticated `
  --cpu 1 `
  --memory 1Gi `
  --port=5000 `
  --project=personal-projects-416300