require_relative 'main'

id = 0

(1..6).each do |id|
  SomeWorker.perform_async(id)
  sleep(10)
end
