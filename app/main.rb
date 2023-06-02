require 'faktory'

class SomeWorker
  include Faktory::Job

  def perform(id)
    puts "Processing #{id}"
    raise StandardError, 'Something went wrong!' if Time.now.strftime("%S").to_i < 30
    puts "Done #{id}"
  end
end
