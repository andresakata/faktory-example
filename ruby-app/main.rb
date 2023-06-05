require 'faktory'

class SomeWorker
  include Faktory::Job
  faktory_options queue: 'ruby-app'

  def perform(id)
    puts "Processing #{id}"
    raise StandardError, 'Ops! Something went wrong.' if Time.now.strftime("%S").to_i < 30
  end
end
