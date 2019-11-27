resource "kubernetes_cron_job" "crathens" {
  metadata {
    name = "crathens"
  }
  spec {
    concurrency_policy            = "Replace"
    failed_jobs_history_limit     = 5
    schedule                      = "0 * * ? * *"
    starting_deadline_seconds     = 10
    successful_jobs_history_limit = 10
    suspend                       = true
    job_template {
      metadata {}
      spec {
        backoff_limit = 2
        template {
          metadata {}
          spec {
            container {
              name  = "crathens"
              image = var.crathens-image-name
            }
            restart_policy = "OnFailure"
          }
        }
      }
    }
  }
}
